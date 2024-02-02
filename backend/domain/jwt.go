package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type JwtData struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
