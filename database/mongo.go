package database

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var dbName string

type Storage interface {
	Create() error
}

type MongoDBStorage struct {
	db *mongo.Client
}

func StartMongoDB() (*MongoDBStorage, error) {

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		return nil, errors.New("you must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	database := os.Getenv("DATABASE")
	if database == "" {
		return nil, errors.New("you must set your 'DATABASE' environmental variable")

	}

	var err error
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		return nil, errors.New("can't verify a connection")
	}

	return &MongoDBStorage{
		db: mongoClient,
	}, nil
}

func CloseMongoDB() {
	err := mongoClient.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
}

func GetCollection(name string) *mongo.Collection {
	return mongoClient.Database(dbName).Collection(name)
}

func (s *MongoDBStorage) Create() error {
	return nil
}
