package main

import (
    "go-api/controller"
	"go-api/db" 
	"go-api/repository"
	"go-api/usecase"
	"github.com/gin-gonic/gin" 
)

func main() {
    server := gin.Default()

    dbConnection, err := db.ConnectDB()
    if err != nil {
        panic(err)
    }

    ProductRepository := repository.NewProductRepository(dbConnection)

    ProductUseCase := usecase.NewProductUsecase(ProductRepository)	

    ProductController := controller.NewProductController(ProductUseCase)

    userRepository := repository.NewUserRepository(dbConnection)
    userUsecase := usecase.NewUserUsecase(userRepository)
    userController := controller.NewUserController(userUsecase)

    // Auth routes
    server.POST("/register", userController.Register)
    server.POST("/login", userController.Login)

    // Products routes
    server.GET("/products", ProductController.GetProducts)
    server.GET("/products/:id", ProductController.GetProductById)
    server.POST("/products", ProductController.CreateProduct)
    server.PUT("/products/:id", ProductController.UpdateProduct)
    server.PATCH("/products/:id", ProductController.PatchProduct)
    server.DELETE("/products/:id", ProductController.DeleteProduct)

    server.Run(":8080")
}