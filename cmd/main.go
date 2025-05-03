package main

import (
	"log"
	"net/http"

	"github.com/alfanzain/keyboard-chiper/internal/core"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
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

	log.Printf("server is listening on %s", ":3066")
	if err := http.ListenAndServe(":3066", r); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to start server: %v", err)
	}
}
