package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CollectionTeam = "teams"

type Team struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Subscribtion string             `json:"subscribtion" bson:"subscription"`
	Owner        string             `json:"owner" bson:"owner"`
	TeamMembers  []string           `json:"teamMembers" bson:"teamMembers"`
}
type TeamUpdate struct {
	Name         string `json:"name" bson:"name"`
	Subscribtion string `json:"subscribtion" bson:"subscription"`
	Owner        string `json:"owner" bson:"owner"`
}
type TeamUsecase interface {
	Create(c context.Context, input Team) error
	GetAll() ([]Team, error)
	GetById(c context.Context, id string) (Team, error)
	Update(id string, input TeamUpdate) error
	Delete(id string) error
	GetAllByOwner(c context.Context, owner string) ([]Team, error)
}
type TeamRepository interface {
	Create(c context.Context, input Team) error
	GetAll() ([]Team, error)
	GetById(c context.Context, id string) (Team, error)
	Update(id string, input TeamUpdate) error
	Delete(id string) error
	GetAllByOwner(c context.Context, owner string) ([]Team, error)
}
