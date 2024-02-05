package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtClaimsAccess struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
type JwtClaimsRefresh struct {
	ID string `json:"id"`
	jwt.StandardClaims
}
type JwtAccessData struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
