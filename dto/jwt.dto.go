package dto

import "github.com/golang-jwt/jwt/v5"

type MyClaims struct {
	jwt.RegisteredClaims
	Name   string
	Email  string
	UserID string
}
