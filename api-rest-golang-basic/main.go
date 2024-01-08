package main

import (
	"github.com/patrik-rangel/API-Rest-Go_CRUD/tree/main/api-rest-golang-basic/configs"
	"github.com/patrik-rangel/API-Rest-Go_CRUD/tree/main/api-rest-golang-basic/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}