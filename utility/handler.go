package utility

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/handlers"
)

// Headers set header to request
func Headers(r http.Handler) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS"})
	return handlers.CORS(headersOk, originsOk, methodsOk)(r)
}

// Response will return json responce of http
// This func handle both error a well as success
func Response(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(payload)
}
