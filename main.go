package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/caarlos0/env/v7"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/mortuie/CRUD_COPY/data_layer"
	"github.com/mortuie/CRUD_COPY/models"
	"github.com/mortuie/CRUD_COPY/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	cfg := models.EnvVariables{}
	err = env.Parse(&cfg)

	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.MongoDbUrl))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	e := routes.ResourceRoutesDeps{Resources: data_layer.ResourceModel{DB: client}}

	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.Logger)

	r.Mount("/{uuid}/{resourceName}", e.ResourceRoutes())

	fmt.Println("Listening on port: ", 3000)
	http.ListenAndServe(":3000", r)
}
