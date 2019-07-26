package jwt

import (
	"go-restapis/config"
	"go-restapis/core/httphandler"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/context"

	"github.com/dgrijalva/jwt-go"
)

// JwtToken , basic jwt model
type JwtToken struct {
	C *config.Configuration
}

// Token ,contains data that will enrypted in JWT token
// When jwt token will decrypt, token model will returns
// Need this model to authenticate and validate resources access by loggedIn user
type Token struct {
	ID string `json:"id"` // User id
	jwt.StandardClaims
}

// CreateToken : takes userId as parameter,
// generates JWT token and
// Return JWT token string
func (jt *JwtToken) CreateToken(id string) (interface{}, error) {

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &Token{
		ID: id,
	})
	// token -> string. Only server knows this secret (foobar).
	tokenString, err := token.SignedString([]byte(jt.C.JwtSecret))
	if err != nil {
		return nil, err
	}
	m := make(map[string]string)
	m["token"] = tokenString // set response data
	return m, nil
}

// ProtectedEndpoint : authenticate all requests
// Takes http handler as params and performs authentication by
// JWT token
// If everythings looks fine, it reqirect the request to actual hadler,
// otherwise, send unauthorized response
func (jt *JwtToken) ProtectedEndpoint(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("middleware", r.URL)
		if strings.Contains(r.URL.Path, "/auth/") {
			h.ServeHTTP(w, r)
		} else {
			// JWT token from request header
			tokenString := r.Header.Get("Authorization")
			// In another way, you can decode token to your struct, which needs to satisfy `jwt.StandardClaims`
			t := Token{}
			token, err := jwt.ParseWithClaims(tokenString, &t, func(token *jwt.Token) (interface{}, error) {
				return []byte(jt.C.JwtSecret), nil
			})
			if !token.Valid || err != nil {
				httphandler.Response(w, httphandler.NewHTTPError(httphandler.Unauthorized, http.StatusUnauthorized))
			} else {
				// Set userId in context, so that we can access it over the request
				context.Set(r, "userId", t.ID)
				// Redirest call to original http handler
				h.ServeHTTP(w, r)
			}
		}

	})
}
