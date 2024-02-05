package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CollectionUser = "users"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Email    string             `json:"email" bson:"email"`
	Role     string             `json:"role" bson:"role"`
}

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
