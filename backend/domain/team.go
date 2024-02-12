package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const TeamCollection = "teams"

type Team struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Subscribtion string             `json:"subscribtion" bson:"subscription"`
	Owner        string             `json:"owner" bson:"owner"`
	TeamMembers  string             `json:"teamMembers" bson:"teamMembers"`
}
type TeamUsecase interface {
	Create()
	GetAll()
	GetById()
	Delete()
	Update()
}
type TeamRepository interface {
	Create()
	GetAll()
	GetById()
	Delete()
	Update()
}
