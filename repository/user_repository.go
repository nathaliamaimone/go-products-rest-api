package repository

import (
    "database/sql"
    "go-api/model"
)

type UserRepository interface {
    CreateUser(user model.User) (model.User, error)
    GetUserByEmail(email string) (model.User, error)
}

type userRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
    return &userRepository{
        db: db,
    }
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
    query := `INSERT INTO users (email, password, role) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
    err := r.db.QueryRow(query, user.Email, user.Password, user.Role).
        Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt)
    return user, err
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
    var user model.User
    query := `SELECT id, email, password, role, created_at, updated_at FROM users WHERE email = $1`
    err := r.db.QueryRow(query, email).
        Scan(&user.Id, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
    return user, err
}