package domain

import "context"

const CollectionUser = "users"

type User struct{}

type UserUpdate struct{}

type UserRepository interface {
	GetById(c context.Context, id string) (*User, error)
	GetByEmail(c context.Context, email string) (*User, error)
	Update(c context.Context, id string, input *UserUpdate) error
	Delete(c context.Context, id string) error
}

type UserUsecase interface {
	GetById(c context.Context, id string) (*User, error)
	GetByEmail(c context.Context, email string) (*User, error)
	Update(c context.Context, id string, input *UserUpdate) error
	Delete(c context.Context, id string) error
}
