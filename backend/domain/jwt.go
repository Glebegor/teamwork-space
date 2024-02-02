package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	RoleId   string `json:"roleId"`
	jwt.StandardClaims
}

type JwtData struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	RoleId   string `json:"roleId"`
}
