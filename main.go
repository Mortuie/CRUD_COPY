package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mortuie/CRUD_COPY/data_layer"
	"github.com/mortuie/CRUD_COPY/routes"
	"github.com/mortuie/CRUD_COPY/utils"
)

func main() {
	cfg := utils.GetEnvVars()
	mStore := utils.SetupMongoDbClient(cfg)

	defer mStore.DB.Disconnect(mStore.CTX)

	e := routes.ResourceRoutesDeps{Resources: data_layer.ResourceModel{DB: mStore.DB}}

	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.Logger)

	r.Mount("/{uuid}/{resourceName}", e.ResourceRoutes())

	fmt.Println("Listening on port: ", 3000)
	http.ListenAndServe(":3000", r)
}
