package repository

import (
	"database/sql"
	"testing"
	"time"
	"go-api/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco de dados: %s", err)
	}
	defer db.Close()


	repo := NewProductRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "created_at", "updated_at"}).
		AddRow(1, "Produto 1", "Descrição 1", 10.5, time.Now(), time.Now()).
		AddRow(2, "Produto 2", "Descrição 2", 20.5, time.Now(), time.Now())

	mock.ExpectQuery("SELECT id, name, description, price, created_at, updated_at FROM product").
		WillReturnRows(rows)

	products, err := repo.GetProducts()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(products))
	assert.Equal(t, 1, products[0].Id)
	assert.Equal(t, "Produto 1", products[0].Name)
	assert.Equal(t, 2, products[1].Id)
	assert.Equal(t, "Produto 2", products[1].Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco de dados: %s", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)
	
	now := time.Now()
	product := model.Product{
		Name:        "Produto Teste",
		Description: "Descrição Teste",
		Price:       15.99,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	mock.ExpectPrepare("INSERT INTO product").
		ExpectQuery().
		WithArgs(product.Name, product.Description, product.Price, product.CreatedAt, product.UpdatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	id, err := repo.CreateProduct(product)
	
	assert.NoError(t, err)	
	assert.Equal(t, 1, id)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestGetProductById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco de dados: %s", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)
	
	now := time.Now()
	
	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "created_at", "updated_at"}).
		AddRow(1, "Produto 1", "Descrição 1", 10.5, now, now)

	mock.ExpectQuery("SELECT id, name, description, price, created_at, updated_at FROM product WHERE id = \\$1").
		WithArgs(1).
		WillReturnRows(rows)

	
	product, err := repo.GetProductById(1)
	

	assert.NoError(t, err)
	assert.Equal(t, 1, product.Id)
	assert.Equal(t, "Produto 1", product.Name)
	assert.Equal(t, "Descrição 1", product.Description)
	assert.Equal(t, 10.5, product.Price)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco de dados: %s", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)
	
	product := model.Product{
		Id:          1,
		Name:        "Produto Atualizado",
		Description: "Descrição Atualizada",
		Price:       25.99,
	}

	mock.ExpectExec("UPDATE product SET name = \\$1, description = \\$2, price = \\$3, updated_at = \\$4 WHERE id = \\$5").
		WithArgs(product.Name, product.Description, product.Price, sqlmock.AnyArg(), product.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.UpdateProduct(product)
	
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco de dados: %s", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	mock.ExpectExec("DELETE FROM product WHERE id = \\$1").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.DeleteProduct(1)
	
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPatchProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco de dados: %s", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)
	
	product := model.Product{
		Id:          1,
		Name:        "Produto Patch",
		Description: "Descrição Patch",
		Price:       30.99,
	}

	mock.ExpectExec("UPDATE product SET name = COALESCE\\(\\$1, name\\), description = COALESCE\\(\\$2, description\\), price = COALESCE\\(\\$3, price\\), updated_at = \\$4 WHERE id = \\$5").
		WithArgs(product.Name, product.Description, product.Price, sqlmock.AnyArg(), product.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.PatchProduct(product)
	
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestErrorHandling(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco de dados: %s", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	mock.ExpectQuery("SELECT id, name, description, price, created_at, updated_at FROM product").
		WillReturnError(sql.ErrConnDone)

	_, err = repo.GetProducts()
	assert.Error(t, err)
	assert.Equal(t, sql.ErrConnDone, err)
	
	mock.ExpectQuery("SELECT id, name, description, price, created_at, updated_at FROM product WHERE id = \\$1").
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	_, err = repo.GetProductById(1)
	assert.Error(t, err)
	assert.Equal(t, sql.ErrNoRows, err)
}