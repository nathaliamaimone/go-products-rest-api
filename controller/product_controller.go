package controller

import (
	"go-api/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

type productController struct {
	//Use case
}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{Id: 1, Name: "Product 1", Price: 10.0},
		{Id: 2, Name: "Product 2", Price: 20.0},
		{Id: 3, Name: "Product 3", Price: 30.0},	
	}
	ctx.JSON(http.StatusOK, products)
}