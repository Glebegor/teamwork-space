package bootstrap

import (
	"context"
	"fmt"
	"time"

	mongoRep "team-work-space/mongo"
	// "go.mongodb.org/mongo-driver/mongo"
)

func NewDBConnection(env *Env) (mongoRep.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%v", env.DBusername, env.DBpassword, env.DBhost, env.DBport)

	client, err := mongoRep.NewClient(mongodbURI)
	if err != nil {
		return client, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return client, err
	}
	err = client.Ping(ctx)
	if err != nil {
		return client, err
	}
	return client, err
}
