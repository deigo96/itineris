package model

import "github.com/golang-jwt/jwt/v5"

type LoginRequest struct {
	NIP      string `json:"nip" validate:"required,numeric,min=18,max=18"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
}

type CustomClaims struct {
	ID       int
	Role     string
	Nip      string
	Register jwt.RegisteredClaims
}
