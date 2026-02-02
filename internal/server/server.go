package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"my-website/internal/config"
	"my-website/internal/handler"
)

type Server struct {
	cfg     *config.Config
	router  *chi.Mux
	handler *handler.Handler
}

func New(cfg *config.Config) *Server {
	return &Server{
		cfg:     cfg,
		router:  chi.NewRouter(),
		handler: handler.New(cfg.TemplDir),
	}
}

func (s *Server) Setup() error {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Compress(5))

	if err := s.handler.LoadTemplates(); err != nil {
		return fmt.Errorf("templates: %w", err)
	}

	fs := http.FileServer(http.Dir(s.cfg.StaticDir))
	s.router.Handle("/static/*", http.StripPrefix("/static/", fs))
	s.router.Get("/", s.handler.Index)

	return nil
}

func (s *Server) Run() error {
	addr := ":" + s.cfg.Port
	log.Printf("Server started at http://localhost%s\n", addr)
	return http.ListenAndServe(addr, s.router)
}
