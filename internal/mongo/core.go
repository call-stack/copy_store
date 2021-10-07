package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCore struct {
}

var Client *mongo.Client

func (m *MongoCore) SetMongoClient() {
	Client, _ = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27018"))

}
