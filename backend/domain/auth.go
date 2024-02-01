package domain

import "context"

const AuthCollection = "users"

type Reg struct{}
type Login struct{}

type AuthRepository interface {
	Create(c context.Context, input Reg) error
	// Login(c context.Context, input Login) error
}
type AuthUsecase interface {
	Register(c context.Context, input Reg) error
	// Login(c context.Context, input Login) error
}
