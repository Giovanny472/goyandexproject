package main

import (
	"log"
	"net/http"

	"github.com/Giovanny472/goyandexproject/internal/app/handlers"
)

func main() {

	http.HandleFunc("/", handlers.FirstIncrement)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
