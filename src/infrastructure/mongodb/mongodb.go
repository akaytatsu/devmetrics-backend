package infra_mongodb

import (
	"app/config"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Database {
	databaseUrl := config.EnvironmentVariables.MONGO_URL
	databaseName := config.EnvironmentVariables.MONGO_DATABASE_NAME

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(databaseUrl))

	if err != nil {
		panic(err)
	}

	return client.Database(databaseName)

}
func ConnectClient() *mongo.Client {
	databaseUrl := config.EnvironmentVariables.MONGO_URL

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(databaseUrl))

	if err != nil {
		panic(err)
	}

	return client

}
