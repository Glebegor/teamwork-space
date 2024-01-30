package bootstrap

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDBConnection(env *Env) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", env.DBhost, env.DBport))
}
