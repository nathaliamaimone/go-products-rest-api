package repository

import (
	"database/sql"
	"testing"
	"time"
	"go-api/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco de dados: %s", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)
	
	now := time.Now()
	user := model.User{
		Email:    "teste@example.com",
		Password: "senha123",
		Role:     "user",
	}

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
		AddRow(1, now, now)

	mock.ExpectQuery("INSERT INTO users \\(email, password, role\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING id, created_at, updated_at").
		WithArgs(user.Email, user.Password, user.Role).
		WillReturnRows(rows)

	createdUser, err := repo.CreateUser(user)
	
	assert.NoError(t, err)
	assert.Equal(t, 1, createdUser.Id)
	assert.Equal(t, "teste@example.com", createdUser.Email)
	assert.Equal(t, "user", createdUser.Role)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco de dados: %s", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)
	
	now := time.Now()
	email := "teste@example.com"
	
	rows := sqlmock.NewRows([]string{"id", "email", "password", "role", "created_at", "updated_at"}).
		AddRow(1, email, "hashed_password", "user", now, now)

	mock.ExpectQuery("SELECT id, email, password, role, created_at, updated_at FROM users WHERE email = \\$1").
		WithArgs(email).
		WillReturnRows(rows)

	user, err := repo.GetUserByEmail(email)
	
	assert.NoError(t, err)
	assert.Equal(t, 1, user.Id)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, "hashed_password", user.Password)
	assert.Equal(t, "user", user.Role)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByEmailNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco de dados: %s", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)
	
	email := "naoexiste@example.com"
	
	mock.ExpectQuery("SELECT id, email, password, role, created_at, updated_at FROM users WHERE email = \\$1").
		WithArgs(email).
		WillReturnError(sql.ErrNoRows)

	_, err = repo.GetUserByEmail(email)

	assert.Error(t, err)
	assert.Equal(t, sql.ErrNoRows, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateUserError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("erro ao criar mock do banco de dados: %s", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)
	
	user := model.User{
		Email:    "teste@example.com",
		Password: "senha123",
		Role:     "user",
	}

	mock.ExpectQuery("INSERT INTO users \\(email, password, role\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING id, created_at, updated_at").
		WithArgs(user.Email, user.Password, user.Role).
		WillReturnError(sql.ErrConnDone)

	_, err = repo.CreateUser(user)

	assert.Error(t, err)
	assert.Equal(t, sql.ErrConnDone, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}