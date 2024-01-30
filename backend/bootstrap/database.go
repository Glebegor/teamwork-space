package bootstrap

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDBConnection(env *Env) (mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", env.DBusername, env.DBpassword, env.DBhost, env.DBport)
	// if env.DBusername == "" || env.DBpassword == "" {
	// mongodbURI = fmt.Sprintf("mongodb://%s:%s", env.DBhost, env.DBport)
	// }
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURI))
	if err != nil {
		return *client, err
	}
}
