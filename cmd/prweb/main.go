package main

import (
	"log"
	"net/http"

	"prweb/internal/handler"

	"github.com/go-chi/chi"
)

func main() {
	h := handler.NewHandler()
	r := chi.NewRouter()
	r.Get("/hello", h.Hello)
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}