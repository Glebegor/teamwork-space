package repository

import (
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

func (uu *userRepository) Create(*domain.User) (*domain.User, error)     { return nil, nil }
func (uu *userRepository) GetById(id string) (*domain.User, error)       { return nil, nil }
func (uu *userRepository) GetByEmail(email string) (*domain.User, error) { return nil, nil }
func (uu *userRepository) Update(*domain.UserUpdate) error               { return nil }
