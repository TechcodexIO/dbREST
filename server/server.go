package server

import (
	"net/http"
	"time"

	"github.com/flarco/g"
	"github.com/labstack/echo/v5"
)

// Server is the main server
type Server struct {
	Port       string
	EchoServer *echo.Echo
	StartTime  time.Time
	Routes     []echo.Route
}

func NewServer() (s *Server) {
	s = &Server{EchoServer: echo.New()}

	// add routes
	for _, route := range standardRoutes {
		s.EchoServer.AddRoute(route)
	}
	return
}

func (s *Server) Start() {
	s.StartTime = time.Now()
	if err := s.EchoServer.Start(":1323"); err != http.ErrServerClosed {
		g.LogFatal(g.Error(err, "could not start server"))
	}
}