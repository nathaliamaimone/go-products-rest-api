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

func (p *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := p.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.Id = productId
	return product, nil
}

func (pu *ProductUsecase) GetProductById(id int) (model.Product, error) {
    return pu.repository.GetProductById(id)
}

func (pu *ProductUsecase) UpdateProduct(product model.Product) error {
    return pu.repository.UpdateProduct(product)
}

func (pu *ProductUsecase) DeleteProduct(id int) error {
    return pu.repository.DeleteProduct(id)
}