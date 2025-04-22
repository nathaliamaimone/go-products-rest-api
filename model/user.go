package model

import "time"

type User struct {
    Id        int       `json:"id"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    Role      string    `json:"role"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}