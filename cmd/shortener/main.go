package main

import (
	"log"
	"net/http"

	"github.com/Giovanny472/goyandexproject/internal/app/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	r.Route("/", handlers.RouterInc)

	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
