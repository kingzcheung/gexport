package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kingzcheung/gexport/internal/handler/api/conn"
	"net/http"
)

type Server struct {
}

func (s *Server) Handle() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/conn", conn.ConnectCtx())
	r.Get("/table/{name}", conn.TableDetailCtx())
	return r
}

func (s *Server) Pattern() string {
	return "/api"
}
