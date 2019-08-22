package core

import (
	"encoding/json"
	"golang-mongodb-restful-starter-kit/config"
	"golang-mongodb-restful-starter-kit/core/httphandler"
	"golang-mongodb-restful-starter-kit/model"
	"golang-mongodb-restful-starter-kit/service/auth"
	"golang-mongodb-restful-starter-kit/service/jwt"
	"golang-mongodb-restful-starter-kit/utility"
	"log"
	"net/http"
)

// AuthHandler
type AuthHadler struct {
	au auth.AuthService
	c  *config.Configuration
}

// NewAuthHandler
func NewAuthHandler(au auth.AuthService, c *config.Configuration) *AuthHadler {
	return &AuthHadler{au, c}

}

// Create
func (h *AuthHadler) Create(w http.ResponseWriter, r *http.Request) {
	requestUser := new(model.User)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)
	requestUser.SetSaltedPassword(requestUser.Password)
	requestUser.Role = utility.UserRole
	err := h.au.Create(r.Context(), requestUser)
	result := make(map[string]interface{})
	if err != nil {
		result = httphandler.NewHTTPError(httphandler.EntityCreationError, http.StatusBadRequest)
	} else {
		result["message"] = "Successfully Registered"
	}
	httphandler.Response(w, result)
}

func (h *AuthHadler) Login(w http.ResponseWriter, r *http.Request) {
	credentials := new(model.Credential)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&credentials)

	user, err := h.au.Login(r.Context(), credentials)
	var result interface{}
	result = make(map[string]interface{})
	if err != nil {
		log.Println(err)
		result = httphandler.NewHTTPError(httphandler.Unauthorized, http.StatusBadRequest)
	} else {
		j := jwt.JwtToken{C: h.c}
		result, err = j.CreateToken(user.ID.Hex(), user.Role)
		if err != nil {
			log.Println(err)
			result = httphandler.NewHTTPError(httphandler.InternalError, 501)
		}
	}
	httphandler.Response(w, result)
}
