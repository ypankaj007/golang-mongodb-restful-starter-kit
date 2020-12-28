package jwt

import (
	"golang-mongodb-restful-starter-kit/config"
	"golang-mongodb-restful-starter-kit/utility"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/context"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtToken , basic jwt model
type JwtToken struct {
	C *config.Configuration
}

// Token ,contains data that will enrypted in JWT token
// When jwt token will decrypt, token model will returns
// Need this model to authenticate and validate resources access by loggedIn user
type Token struct {
	ID   string `json:"id"`   // User id
	Role string `json:"role"` // user role
	jwt.StandardClaims
}

// CreateToken : takes userId as parameter,
// generates JWT token and
// Return JWT token string
func (jt *JwtToken) CreateToken(id, role string) (map[string]string, error) {

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &Token{
		ID:   id,
		Role: role,
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
		if strings.Contains(r.URL.Path, "/auth/") || strings.Contains(r.URL.Path, "/swagger/") {
			h.ServeHTTP(w, r)
		} else {
			// JWT token from request header
			tokenString := r.Header.Get("Authorization")
			// In another way, you can decode token to your struct, which needs to satisfy `jwt.StandardClaims`
			t := Token{}
			token, err := jwt.ParseWithClaims(tokenString, &t, func(token *jwt.Token) (interface{}, error) {
				return []byte(jt.C.JwtSecret), nil
			})
			if err != nil || !token.Valid {
				utility.Response(w, utility.NewHTTPError(utility.Unauthorized, http.StatusUnauthorized))
			} else {

				// Set userId and in context, so that we can access it over the request
				// is some request, we need loggedIn user information
				context.Set(r, "userId", t.ID) // set loggedIn user id in context
				context.Set(r, "role", t.Role) // set loggedIn user role in context
				// Redirest call to original http handler
				h.ServeHTTP(w, r)
			}
		}

	})
}
