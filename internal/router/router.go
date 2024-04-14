package router

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pdstuber/go-echo-test/internal/api"
	"github.com/pdstuber/go-echo-test/pkg/jsonserializer"
)

type Server struct {
	listenPort string
	errChan    chan error
	echo       *echo.Echo
}

func New(listenPort string) *Server {
	e := echo.New()
	e.JSONSerializer = jsonserializer.NewSerializer()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api", api.Hello)

	return &Server{
		listenPort: listenPort,
		echo:       e,
	}
}

func (s *Server) Start(ctx context.Context) error {
	go func() {
		if err := s.echo.Start(s.listenPort); err != nil {
			s.errChan <- err
		}
	}()

	select {
	case err := <-s.errChan:
		return err
	case <-ctx.Done():
		return nil
	}
}

func (s *Server) Stop(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}
