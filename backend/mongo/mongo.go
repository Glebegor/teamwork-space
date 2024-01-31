package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database interface {
	Collections(string) Collection
	Client() Client
}

type Collection interface {
}

type SingleResult interface {
}

type Cursor interface {
}
type Client interface {
	Connect(context.Context) error
	Ping(context.Context) error
}

// Structures
type mongoClient struct {
	cl *mongo.Client
}

type mongoCursor struct {
}

type mongoDatabase struct {
}
type mongoCollection struct {
}

type mongoSingleResult struct {
}

type mongoSession struct {
}

type nullwareDecoder struct {
}

func NewClient(connection string) (Client, error) {
	time.Local = time.UTC
	c, err := mongo.NewClient(options.Client().ApplyURI(connection))
	return &mongoClient{cl: c}, err
}
func (mc *mongoClient) Connect(ctx context.Context) error {
	return mc.cl.Connect(ctx)
}
func (mc *mongoClient) Ping(ctx context.Context) error {
	return mc.cl.Ping(ctx, readpref.Primary())
}
