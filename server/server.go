package server

import (
	"errors"
	"fmt"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/usamaiqbal83/leadfuzz-test/domain"
	"gopkg.in/tylerb/graceful.v1"
)

// Config type
type Config struct {
}

// Options for running the server
type Options struct {
	// if value is 0, it means that server will wait for all active requests to finish before shut down
	Timeout time.Duration
	// handler to call before shutting down
	ShutdownHandler func()
}

// options for CORS configurations
type CORSOptions struct {
	AllowOrigin  []string
	AllowMethods []string
}

// Server type
type Server struct {
	echo           *echo.Echo
	gracefulServer *graceful.Server
	timeout        time.Duration
}

func NewServer(options *Config) *Server {

	// setup server
	e := echo.New()

	s := &Server{echo: e, gracefulServer: nil}

	return s
}

func (s *Server) Run(address string, options *Options) {
	s.timeout = options.Timeout

	// assign listening address
	s.echo.Server.Addr = address

	s.gracefulServer = &graceful.Server{
		Timeout:           options.Timeout,
		Server:            s.echo.Server,
		ShutdownInitiated: options.ShutdownHandler,
	}

	// gracefully server
	graceful.ListenAndServe(s.echo.Server, options.Timeout)
}

func (s *Server) Stop() {
	s.gracefulServer.Stop(s.timeout)
}

func (s *Server) ConfigureCORS(options *CORSOptions) {
	s.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: options.AllowOrigin,
		AllowMethods: options.AllowMethods,
	}))
}

func (s *Server) AddResources(resources ...domain.IResource) {
	for _, resource := range resources {
		if resource.Routes() == nil {
			// server/router instantiation error
			// its safe to throw panic here
			panic(errors.New(fmt.Sprintf("Routes definition missing: %v", resource)))
		}

		s.addRoutes(resource.Routes())
	}
}

func (s *Server) addRoutes(routes *domain.Routes) {
	if routes == nil {
		return
	}

	// add all routes
	for _, route := range *routes {
		s.echo.Add(route.Method, route.Pattern, route.RouteHandler)
	}
}
