package repository

import (
	"context"
	"team-work-space/domain"
	"team-work-space/mongo"
)

type authRepository struct {
	database   mongo.Database
	collection string
}

func NewAuthRepository(database mongo.Database, collection string) domain.AuthRepository {
	return &authRepository{
		database:   database,
		collection: collection,
	}
}

func (ar *authRepository) Create(c context.Context, input domain.Reg) error {
	collection := ar.database.Collection(ar.collection)
	_, err := collection.InsertOne(c, input)
	return nil
}
