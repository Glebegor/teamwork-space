package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"team-work-space/domain"
	"team-work-space/mongo"
)

type teamRepository struct {
	database   mongo.Database
	collection string
}

func NewTeamRepository(database mongo.Database, collection string) domain.TeamRepository {
	return &teamRepository{
		database:   database,
		collection: collection,
	}
}

func (tr *teamRepository) Create(c context.Context, input domain.Team) error {
	_, err := tr.database.Collection(tr.collection).InsertOne(c, input)
	return err
}
func (tr *teamRepository) GetAll() ([]domain.Team, error) {
	return nil, nil
}
func (tr *teamRepository) GetById(c context.Context, id string) (domain.Team, error) {
	var team domain.Team
	err := tr.database.Collection(tr.collection).FindOne(c, bson.M{"_id": id}).Decode(&team)
	if err != nil {
		return team, err
	}
	return team, nil
}
func (tr *teamRepository) Update(id string, input domain.TeamUpdate) error {
	return nil
}
func (tr *teamRepository) Delete(id string) error {
	return nil
}
func (tu *teamUsecase) GetAllByOwner(c context.Context, owner string) ([]Team, error) {
	return nil, nil
}
