package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alfanzain/go-keyboard-chiper/internal/core"
	"github.com/caarlos0/env/v11"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	s, err := core.NewService(core.ServiceConfig{})
	if err != nil {
		log.Fatalf("failed to initialize service: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/decode", func(w http.ResponseWriter, r *http.Request) {
		output, err := s.HandleDecode(r.Context(), r.URL.Query().Get("input"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, map[string]string{"output": output})
	})

	log.Printf("server is listening on %s", cfg.ListenPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.ListenPort), r); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to start server: %v", err)
	}
}

type config struct {
	ListenPort string `env:"LISTEN_PORT,required"`
}
