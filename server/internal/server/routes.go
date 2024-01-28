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

	// Handle NoAuth Routes
	noAuth := apiRouter.PathPrefix("").Subrouter()
	{
		noAuth.Handle("/health", mfehandlers.HealthCheck()).Methods(http.MethodGet)
		noAuth.Handle("/health", mfehandlers.PostHealthCheck()).Methods(http.MethodPost)
	}
	noAuth = appShellRouter.PathPrefix("").Subrouter()
	{
		noAuth.Handle("/", mfehandlers.RenderAppShells()).Methods("GET")
	}

	// TODO: Handle Auth Routes

	// Setup middlewares
	s.router.Use(middleware.Logger(os.Stdout))
}
