package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct { 
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (p *ProductUsecase) GetProducts() ([]model.Product, error) {
	return p.repository.GetProducts()
}