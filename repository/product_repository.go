package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connention *sql.DB
}

func NewProductRepository(connention *sql.DB) ProductRepository {
	return ProductRepository{
		connention: connention,
	}	
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
    query := "SELECT id, name, description, price, created_at, updated_at FROM product"
    rows, err := pr.connention.Query(query)
    if err != nil {
        fmt.Println(err)
        return []model.Product{}, err
    }

    var productList []model.Product
    var productObject model.Product

    for rows.Next() {
        err = rows.Scan(
            &productObject.Id,
            &productObject.Name,
            &productObject.Description,
            &productObject.Price,
            &productObject.CreatedAt,
            &productObject.UpdatedAt)

        if err != nil {
            fmt.Println(err)
            return []model.Product{}, err    
        }

        productList = append(productList, productObject)
    }

    rows.Close()
    return productList, nil
}