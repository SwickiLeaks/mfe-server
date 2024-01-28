package mfehandlers

import (
	"mfeserver/internal/utils"
	"net/http"
)

// Convenient typing
type M map[string]interface{}

// Server Health Check endpoint
func HealthCheck() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		resp := M{
			"status":  "available",
			"message": "healthy",
			"data":    M{"hello": "beautiful"},
		}
		utils.WriteJSON(rw, http.StatusOK, resp)
	})
}

// Server Health Check endpoint
func PostHealthCheck() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		resp := M{
			"status":  "available",
			"message": "healthy",
			"data":    M{"hello": "post"},
		}
		utils.WriteJSON(rw, http.StatusOK, resp)
	})
}

// Render appshells
func RenderAppShells() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		resp := M{
			"status":  "available",
			"message": "healthy",
			"data":    M{"render": "appshells"},
		}
		utils.WriteJSON(rw, http.StatusOK, resp)
	})
}
