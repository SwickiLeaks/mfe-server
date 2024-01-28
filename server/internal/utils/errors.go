package utils

import (
	"log"
	"net/http"
)

// Convenient typing
type M map[string]interface{}

// Format Server Errors
func ServerError(w http.ResponseWriter, err error) {
	log.Println(err)
	ErrorResponse(w, http.StatusInternalServerError, "internal error")
}

// Produce Error Response
func ErrorResponse(w http.ResponseWriter, code int, errs interface{}) {
	WriteJSON(w, code, M{"errors": errs})
}
