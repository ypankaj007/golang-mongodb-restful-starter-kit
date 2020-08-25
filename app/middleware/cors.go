package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func Cors(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}
