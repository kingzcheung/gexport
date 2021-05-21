package internal

import (
	"context"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/kingzcheung/gexport/internal/handler"
	"github.com/kingzcheung/gexport/internal/handler/api"
	"github.com/kingzcheung/gexport/internal/handler/web"
	"golang.org/x/sync/errgroup"
	"net/http"
)

type Server struct {
	Handler *chi.Mux
	Session *scs.SessionManager
}

func NewServer(session *scs.SessionManager) *Server {
	handle := provideRouter(
		&api.Server{},
		&web.Server{},
	)
	return &Server{
		Handler: handle,
		Session: session,
	}
}

func (s Server) ListenAndServe(ctx context.Context, addr string) error {
	var g errgroup.Group
	s1 := &http.Server{
		Addr:    addr,
		Handler: s.Session.LoadAndSave(s.Handler),
	}
	g.Go(func() error {
		select {
		case <-ctx.Done():
			return s1.Shutdown(ctx)
		}
	})
	g.Go(func() error {
		return s1.ListenAndServe()
	})
	return g.Wait()
}

func provideRouter(hs ...handler.Handler) *chi.Mux {
	r := chi.NewRouter()
	for _, h := range hs {
		r.Mount(h.Pattern(), h.Handle())
	}
	return r
}
