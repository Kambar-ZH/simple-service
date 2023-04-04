package dtos

import (
	"github.com/Kambar-ZH/simple-service/internal/tools/auth_tool"
)

type RegisterRequest struct {
	Email    string             `gorm:"email" json:"email" binding:"required,email"`
	Password auth_tool.Password `gorm:"password" json:"password" binding:"required,min=8,max=20"`
}

type RegisterResponse struct {
	ID int64 `gorm:"id" json:"id"`
}

type TokenPair struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type LoginRequest struct {
	Email    string             `json:"email" binding:"required,email"`
	Password auth_tool.Password `json:"password" binding:"required,min=8,max=20"`
}

type LoginResponse struct {
	TokenPair `json:"tokens"`
}

type RefreshRequest struct {
	Refresh string `json:"refresh" binding:"required"`
}

type RefreshResponse struct {
	TokenPair `json:"tokens"`
}
