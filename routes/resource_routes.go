package routes

import (
	"encoding/json"
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
	r.Get("/{resourceId}", deps.get_one)
	r.Post("/", deps.create)
	r.Delete("/{resourceId}", deps.delete)

	return r
}

func write_err_response(w http.ResponseWriter, message string, statusCode int) {
	b, _ := json.Marshal(models.ErrorResponse{Message: message})
	w.WriteHeader(statusCode)
	w.Write(b)
}

func (deps ResourceRoutesDeps) get_one(w http.ResponseWriter, r *http.Request) {
	rp := r.Context().Value("request_params").(models.RequestParams)
	resourceId := chi.URLParam(r, "resourceId")

	if resourceId == "" {
		write_err_response(w, "ResourceId not valid", 400)
		return
	}

	res, err := deps.Resources.GetResource(rp.Uuid+rp.Resource, resourceId)

	if err != nil {
		write_err_response(w, err.Error(), 400)
		return
	}

	if len(res) == 0 {
		write_err_response(w, "Resource not found", 404)
		return
	}

	b, _ := json.Marshal(res)
	w.WriteHeader(200)
	w.Write(b)
}

func (deps ResourceRoutesDeps) list(w http.ResponseWriter, r *http.Request) {
	rp := r.Context().Value("request_params").(models.RequestParams)

	res, err := deps.Resources.GetResources(rp.Uuid + rp.Resource)

	if err != nil {
		write_err_response(w, "Error listing resources", 404)
		return
	}

	b, _ := json.Marshal(res)

	w.WriteHeader(200)
	w.Write(b)
}

func (deps ResourceRoutesDeps) create(w http.ResponseWriter, r *http.Request) {
	rp := r.Context().Value("request_params").(models.RequestParams)
	deps.Resources.CreateCollection(rp.Uuid + rp.Resource)

	var anyJson map[string]interface{}

	json.NewDecoder(r.Body).Decode(&anyJson)
	delete(anyJson, "_id")

	if anyJson == nil {
		write_err_response(w, "JSON body is empty.", 400)
		return
	}

	res, err := deps.Resources.InsertResource(rp.Uuid+rp.Resource, anyJson)

	if err != nil {
		write_err_response(w, "Internal server error", 500)
		return
	}

	b, _ := json.Marshal(res)
	w.WriteHeader(200)
	w.Write(b)
}

func (deps ResourceRoutesDeps) update(w http.ResponseWriter, r *http.Request) {
	rp := r.Context().Value("request_params").(models.RequestParams)

	_ = rp
	w.Write([]byte{})
}

func (deps ResourceRoutesDeps) delete(w http.ResponseWriter, r *http.Request) {
	rp := r.Context().Value("request_params").(models.RequestParams)
	resourceId := chi.URLParam(r, "resourceId")

	if resourceId == "" {
		write_err_response(w, "ResourceId not valid", 400)
		return
	}

	err := deps.Resources.DeleteResource(rp.Uuid+rp.Resource, resourceId)

	if err != nil {
		write_err_response(w, err.Error(), 500)
	}

	w.WriteHeader(200)
	w.Write([]byte{})
}
