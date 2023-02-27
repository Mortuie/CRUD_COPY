package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type crudResource struct{}

func (rs crudResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Use(middlewareTest)
	r.Get("/", rs.List)

	return r
}

func (rs crudResource) List(w http.ResponseWriter, r *http.Request) {
	rp := r.Context().Value("request_params").(RequestParams)

	fmt.Printf("%+v\n", rp)

	w.Write([]byte("memes"))
}
