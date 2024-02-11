package bootstrap

import (
	"github.com/sirupsen/logrus"
	// "go.mongodb.org/mongo-driver/mongo"
	"team-work-space/mongo"
)

type Application struct {
	Mongo mongo.Client
	Env   *Env
}

func App() Application {
	app := &Application{}
	env, err := NewEnv()
	if err != nil {
		logrus.Fatalf("Error while getting environment variables: %v", err.Error())
	}
	app.Env = env
	mongo, err := NewDBConnection(env)
	if err != nil {
		logrus.Fatalf("Error while connected to database: %v", err.Error())
	}
	app.Mongo = mongo
	return *app
}
