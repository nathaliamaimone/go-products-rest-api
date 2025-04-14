package main

import (
	"go-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
    server := gin.Default()

    ProductController := controller.NewProductController()
    
    server.GET("/products", ProductController.GetProducts)

    server.Run(":8080")
}