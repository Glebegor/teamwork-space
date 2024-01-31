package bootstrap

import (
	"log"

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
		log.Fatal(err.Error())
	}
	app.Env = env
	mongo, err := NewDBConnection(env)
	if err != nil {
		log.Fatal(err.Error())
	}
	app.Mongo = mongo
	return *app
}
