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

type JwtToken struct {
	C *config.Configuration
}

type Token struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

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
	m["token"] = tokenString
	return m, nil
}

func (jt *JwtToken) ProtectedEndpoint(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("middleware", r.URL)
		if strings.Contains(r.URL.Path, "/auth/") {
			h.ServeHTTP(w, r)
		} else {
			tokenString := r.Header.Get("Authorization")
			// In another way, you can decode token to your struct, which needs to satisfy `jwt.StandardClaims`
			t := Token{}
			token, err := jwt.ParseWithClaims(tokenString, &t, func(token *jwt.Token) (interface{}, error) {
				return []byte(jt.C.JwtSecret), nil
			})
			if !token.Valid || err != nil {
				httphandler.Response(w, httphandler.NewHttpError(httphandler.Unauthorized, http.StatusUnauthorized))
			} else {
				context.Set(r, "userId", t.ID)
				h.ServeHTTP(w, r)
			}
		}

	})
}
