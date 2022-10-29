package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/matheusF23/crud-go-postgres/configs"
	"github.com/matheusF23/crud-go-postgres/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	router := chi.NewRouter()
	router.Post("/", handlers.Create)
	router.Put("/{id}", handlers.Update)
	router.Delete("/{id}", handlers.Delete)
	router.Get("/", handlers.List)
	router.Get("/{id}", handlers.Get)
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router)
}
