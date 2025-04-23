package main

import (
	"go-api/config"
	"go-api/controller"
	"go-api/db"
	"go-api/middleware"
	"go-api/repository"
	"go-api/service"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
    if err := config.LoadConfig(); err != nil {
        panic(err)
    }
    
    server := gin.Default()

    dbConnection, err := db.ConnectDB()
    if err != nil {
        panic(err)
    }

    ProductRepository := repository.NewProductRepository(dbConnection)

    ProductUseCase := usecase.NewProductUsecase(ProductRepository)	

    ProductController := controller.NewProductController(ProductUseCase)

    jwtService := service.NewJWTService()

    userRepository := repository.NewUserRepository(dbConnection)

    userUsecase := usecase.NewUserUsecase(userRepository)

    userController := controller.NewUserController(userUsecase, jwtService)

    // Public routes
    server.POST("/register", userController.Register)
    server.POST("/login", userController.Login)
    server.GET("/products", ProductController.GetProducts)
    server.GET("/products/:id", ProductController.GetProductById)

    // Protected routes
    protected := server.Group("/")
    protected.Use(middleware.AuthMiddleware(jwtService))
    {
        protected.POST("/products", ProductController.CreateProduct)
        protected.PUT("/products/:id", ProductController.UpdateProduct)
        protected.PATCH("/products/:id", ProductController.PatchProduct)
        protected.DELETE("/products/:id", ProductController.DeleteProduct)
    }

    server.Run(":8080")
}