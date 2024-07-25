package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/israelalvesmelo/desafio-multithreading/internal/infra/webserver/handlers"
)

func main() {
	cepHandler := handlers.NewCepHandler()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/cep/{cep}", cepHandler.GetCep)

	http.ListenAndServe(":8080", router)
}
