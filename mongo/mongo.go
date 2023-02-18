package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(constant.MongoUrl))
	if err != nil {
		panic("error mongo client")
	}
	return client
}
