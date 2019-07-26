package core

import (
	"encoding/json"
	"go-restapis/config"
	"go-restapis/core/httphandler"
	"go-restapis/model"
	"go-restapis/service/auth"
	"go-restapis/service/jwt"
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

func (h *AuthHadler) Create(w http.ResponseWriter, r *http.Request) {
	requestUser := new(model.User)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)
	requestUser.SetSaltedPassword(requestUser.Password)
	err := h.au.Create(r.Context(), requestUser)
	result := make(map[string]interface{})
	if err != nil {
		result = httphandler.NewHttpError(httphandler.EntityCreationError, http.StatusBadRequest)
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
		result = httphandler.NewHttpError(httphandler.Unauthorized, http.StatusBadRequest)
	} else {
		j := jwt.JwtToken{C: h.c}
		result, err = j.CreateToken(user.ID.Hex())
		if err != nil {
			log.Println(err)
			result = httphandler.NewHttpError(httphandler.InternalError, 501)
		}
	}
	httphandler.Response(w, result)
}
