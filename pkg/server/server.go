package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Config struct {
	Port int `env:"PORT" envDefault:"8080"`
}

type Server struct {
	cfg       Config
	cancelApp context.CancelCauseFunc

	http *http.Server
	gin  *gin.Engine
}

//nolint:mnd
func New(
	cfg Config,
	cancelApp context.CancelCauseFunc,
) (*Server, error) {
	res := &Server{
		cfg:       cfg,
		cancelApp: cancelApp,
		gin:       gin.New(),
	}

	res.http = &http.Server{
		Addr:              fmt.Sprintf("127.0.0.1:%d", res.cfg.Port),
		Handler:           res.gin,
		ReadHeaderTimeout: 10 * time.Second,
	}

	res.setupMiddleware()

	return res, nil
}

func (s *Server) setupMiddleware() {
	s.gin.Use(gin.Recovery())
}

func (s *Server) OnStart() {
	go func() {
		err := s.http.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.cancelApp(err)
		} else {
			s.cancelApp(nil)
		}
	}()
}

func (s *Server) OnStop(ctx context.Context) error {
	err := s.http.Shutdown(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *Server) GetEngine() *gin.Engine {
	return s.gin
}

func (s *Server) GetRouter() gin.IRouter {
	return s.gin
}
