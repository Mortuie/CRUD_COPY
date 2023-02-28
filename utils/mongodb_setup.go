package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mortuie/CRUD_COPY/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var databases = []string{"resources"}

type MongoStore struct {
	DB  *mongo.Client
	CTX context.Context
}

func SetupMongoDbClient(cfg models.EnvVariables) *MongoStore {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.MongoDbUrl))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if err != nil {
		log.Fatal(err)
	}

	client.Connect(ctx)

	setupMongoDatabases(client)

	return &MongoStore{DB: client, CTX: ctx}
}

func setupMongoDatabases(c *mongo.Client) {
	for _, database_name := range databases {
		fmt.Println("creating database:", database_name)
		c.Database(database_name)
	}
}
