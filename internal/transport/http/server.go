package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/render"
)

type Config struct {
	Port         int           `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *Config, handler http.Handler) (*Server, error) {
	return &Server{
		httpServer: &http.Server{
			ReadHeaderTimeout: cfg.ReadTimeout,
			Addr:              fmt.Sprintf(":%d", cfg.Port),
			Handler:           handler,
			ReadTimeout:       cfg.ReadTimeout,
			WriteTimeout:      cfg.WriteTimeout,
		},
	}, nil
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func ResponseJSON(w http.ResponseWriter, r *http.Request, obj interface{}) {
	if obj == nil {
		obj = struct {
		}{}
	}
	render.JSON(w, r, obj)
}

func routerCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("x-force", "with-you")
	ResponseJSON(w, r, nil)
}
