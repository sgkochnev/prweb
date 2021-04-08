package main

import (
	"log"
	"net/http"

	"prweb/internal/config"
	"prweb/internal/handler"

	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	cfg := config.Srever{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler()

	r := chi.NewRouter()
	r.Get("/hello", h.Hello)

	path := cfg.Host + ":" + cfg.Port

	log.Printf("starting server at %s", path)
	err = http.ListenAndServe(path, r)
	log.Fatal(err)
	log.Print("shutting server down")
}
