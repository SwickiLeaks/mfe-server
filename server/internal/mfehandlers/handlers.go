package mfehandlers

import (
	"mfeserver/internal/utils"
	"net/http"
	"path/filepath"
)

// Convenient typing
type M map[string]interface{}

// Server Health Check endpoint
func HealthCheck() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		resp := M{
			"status":  "available",
			"message": "healthy",
			"data":    M{"hello": "updated"},
		}
		utils.WriteJSON(rw, http.StatusOK, resp)
	})
}

// Render appshells
func RenderAppShells() http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		staticDir := "./static"
		indexPath := filepath.Join(staticDir, "index.html")
		http.ServeFile(rw, r, indexPath)
	})
}
