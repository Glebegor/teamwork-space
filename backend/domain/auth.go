package domain

import (
	"context"
)

const AuthCollection = "users"

type Reg struct {
	Username string `form:"username" json:"username"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}
type Login struct{}

type AuthRepository interface {
	Create(c context.Context, input User) error
	GetByEmail(c context.Context, email string) (User, error)
	// Login(c context.Context, input Login) error
}
type AuthUsecase interface {
	GetByEmail(c context.Context, email string) (User, error)
	Register(c context.Context, input User) error
	// Login(c context.Context, input Login) error
}
