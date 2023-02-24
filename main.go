package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Mount("/{uuid}/{resourceName}", crudResource{}.Routes())

	fmt.Println("Listening on port: ", 3000)
	http.ListenAndServe(":3000", r)
}
