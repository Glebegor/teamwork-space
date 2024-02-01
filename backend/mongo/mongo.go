package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database interface {
	Collection(string) Collection
	Client() Client
}

type Collection interface {
	InsertOne(ctx context.Context, data interface{}) (interface{}, error)
	FindOne(c context.Context, filter interface{}) SingleResult
}

type SingleResult interface {
	Decode(v interface{}) error
}

type Cursor interface {
	Decode(v interface{}) error
}
type Client interface {
	Database(string) Database
	Connect(context.Context) error
	Ping(context.Context) error
}

// Structures
type mongoClient struct {
	cl *mongo.Client
}

type mongoCursor struct {
	mc *mongo.Cursor
}

type mongoDatabase struct {
	db *mongo.Database
}
type mongoCollection struct {
	coll *mongo.Collection
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

type mongoSession struct {
	mongo.Session
}

type nullwareDecoder struct {
}

func (mc *mongoCollection) InsertOne(ctx context.Context, data interface{}) (interface{}, error) {
	id, err := mc.coll.InsertOne(ctx, data)
	return id.InsertedID, err
}
func (mc *mongoCollection) FindOne(c context.Context, filter interface{}) SingleResult {
	SingleResult := mc.coll.FindOne(c, filter)
	return &mongoSingleResult{sr: SingleResult}
}

func (md *mongoDatabase) Collection(colName string) Collection {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

func (md *mongoDatabase) Client() Client {
	client := md.db.Client()
	return &mongoClient{cl: client}
}

func (mc *mongoClient) Database(dbName string) Database {
	db := mc.cl.Database(dbName)
	return &mongoDatabase{db: db}
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

func (ms *mongoSingleResult) Decode(v interface{}) error {
	return ms.sr.Decode(v)
}

func (mr *mongoCursor) Decode(v interface{}) error {
	return mr.mc.Decode(v)
}
