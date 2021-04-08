package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"prweb/internal/api/jokes"
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

	apiClient := jokes.NewJokeClient(cfg.JokeURL)

	h := handler.NewHandler(apiClient, cfg.CustomJoke)

	r := chi.NewRouter()
	r.Get("/hello", h.Hello)

	path := cfg.Host + ":" + cfg.Port

	srv := &http.Server{
		Addr:    path,
		Handler: r,
	}

	// handler shutdown gracefully
	quit := make(chan os.Signal, 1)
	done := make(chan error, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		err := srv.Shutdown(ctx)
		done <- err
	}()

	log.Printf("starting server at %s", path)
	_ = srv.ListenAndServe()

	err = <-done
	log.Printf("shutting server down with %v", err)
}
