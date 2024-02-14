package repository

import (
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

func (tr *teamRepository) Create(input domain.Team) error {
	return nil
}
func (tr *teamRepository) GetAll() ([]domain.Team, error) {
	return nil, nil
}
func (tr *teamRepository) GetById(id string) (*domain.Team, error) {
	return nil, nil
}
func (tr *teamRepository) Update(id string, input *domain.TeamUpdate) error {
	return nil
}
func (tr *teamRepository) Delete(id string) error {
	return nil
}
