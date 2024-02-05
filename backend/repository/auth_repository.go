package repository

import (
	"context"
	"fmt"
	"team-work-space/domain"
	"team-work-space/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (ar *authRepository) Create(c context.Context, input domain.User) error {
	collection := ar.database.Collection(ar.collection)
	id, err := collection.InsertOne(c, input)
	fmt.Print(id)
	return err
}
func (ar *authRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	collection := ar.database.Collection(ar.collection)
	var user domain.User

	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}
func (ar *authRepository) GetById(c context.Context, id string) (domain.User, error) {
	var user domain.User
	collection := ar.database.Collection(ar.collection)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}
