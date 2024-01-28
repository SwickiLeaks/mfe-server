package middleware

import (
	"io"
	"net/http"

	"github.com/gorilla/handlers"
)

// Middleware for Logging
// TODO: Enhance logging experience
func Logger(w io.Writer) func(h http.Handler) http.Handler {
	return (func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(w, h)
	})
}
