package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type crudResource struct{}

func (rs crudResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List)

	return r
}

func (rs crudResource) List(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	rsName := chi.URLParam(r, "resourceName")
	fmt.Println(uuid, rsName)
	w.Write([]byte("memes"))
}
