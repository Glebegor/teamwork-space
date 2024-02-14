package repository

import (
	"context"
	"team-work-space/domain"
	"team-work-space/mongo"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(database mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   database,
		collection: collection,
	}
}
func (ur *userRepository) GetById(c context.Context, id string) (*domain.User, error) {
	return nil, nil
}
func (ur *userRepository) GetByEmail(c context.Context, email string) (*domain.User, error) {
	return nil, nil
}
func (ur *userRepository) Update(c context.Context, id string, input *domain.UserUpdate) error {
	return nil
}
func (ur *userRepository) Delete(c context.Context, id string) error { return nil }
