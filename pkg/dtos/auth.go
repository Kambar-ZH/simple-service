package dtos

import "github.com/golang-jwt/jwt/v4"

type RegisterRequest struct {
	Email    string `gorm:"email" json:"email"`
	Password string `gorm:"password" json:"password"`
}

type RegisterResponse struct {
	ID int64 `gorm:"id" json:"id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
