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
func (uu *userRepository) GetById(c context.Context, id string) (*domain.User, error) {
	return nil, nil
}
func (uu *userRepository) GetByEmail(c context.Context, email string) (*domain.User, error) {
	return nil, nil
}
func (uu *userRepository) Update(c context.Context, id string, input *domain.UserUpdate) error {
	return nil
}
func (uu *userRepository) Delete(c context.Context, id string) error { return nil }
