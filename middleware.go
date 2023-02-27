package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func middlewareTest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uuid := chi.URLParam(r, "uuid")
		sourceName := chi.URLParam(r, "resourceName")

		rp := &RequestParams{Uuid: uuid, SourceName: sourceName}
		err := validator.New().Struct(rp)

		if err != nil {
			b, _ := json.Marshal(&ErrorResponse{Message: "query params fail validation"})
			w.WriteHeader(400)
			w.Write(b)
			return
		}

		ctx := context.WithValue(r.Context(), "request_params", *rp)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
