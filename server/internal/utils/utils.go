package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// Write a JSON API response
func WriteJSON(w http.ResponseWriter, code int, data interface{}) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(jsonBytes)

	if err != nil {
		log.Println(err)
	}
}
