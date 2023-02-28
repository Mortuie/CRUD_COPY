package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mortuie/CRUD_COPY/data_layer"
	"github.com/mortuie/CRUD_COPY/middleware"
	"github.com/mortuie/CRUD_COPY/models"
)

type ResourceRoutesDeps struct {
	Resources data_layer.ResourceModel
}

func (deps ResourceRoutesDeps) ResourceRoutes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Resource_request_param_validation)
	r.Get("/", deps.list)
	r.Post("/", deps.create)

	return r
}

func (deps ResourceRoutesDeps) list(w http.ResponseWriter, r *http.Request) {
	rp := r.Context().Value("request_params").(models.RequestParams)

	fmt.Printf("%+v\n", rp)

	w.Write([]byte("memes"))
}

func (deps ResourceRoutesDeps) create(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("creating..."))
}
