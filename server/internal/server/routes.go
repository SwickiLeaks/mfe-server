package server

import (
	"mfeserver/internal/mfehandlers"
	"mfeserver/internal/middleware"
	"net/http"
	"os"
)

const (
	MustAuth     = true
	OptionalAuth = true
)

// Define and configure ALL server routes and middleware
func (s *Server) routes() {
	appShellRouter := s.router.PathPrefix("").Subrouter()
	apiRouter := s.router.PathPrefix("/api/v1").Subrouter()

	// Handle NoAuth API Routes
	noAuth := apiRouter.PathPrefix("").Subrouter()
	{
		noAuth.Handle("/health", mfehandlers.HealthCheck()).Methods(http.MethodGet)
	}

	// Handle Appshell Content
	noAuth = appShellRouter.PathPrefix("").Subrouter()
	{
		noAuth.Handle("/", mfehandlers.RenderAppShells()).Methods(http.MethodGet)
	}

	// Setup middlewares
	// NOTE: Order important
	s.router.Use(middleware.Logger(os.Stdout))
	s.router.Use(middleware.Middleware1())
	s.router.Use(middleware.Middleware2())
}
