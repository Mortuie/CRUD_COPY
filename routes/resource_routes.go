package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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

	deps.Resources.CreateCollection("memes")

	fmt.Printf("%+v\n", rp)

	w.Write([]byte("memes"))
}

func (deps ResourceRoutesDeps) create(w http.ResponseWriter, r *http.Request) {
	rp := r.Context().Value("request_params").(models.RequestParams)

	_ = rp

	var anyJson map[string]interface{}

	json.NewDecoder(r.Body).Decode(&anyJson)

	if anyJson == nil {
		b, _ := json.Marshal(models.ErrorResponse{Message: "JSON body is empty."})
		w.WriteHeader(400)
		w.Write(b)
		return
	}

	anyJson["id"] = uuid.New().String()

	w.Write([]byte("creating..."))
}
