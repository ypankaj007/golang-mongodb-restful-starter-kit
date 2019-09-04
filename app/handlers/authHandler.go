package handlers

import (
	"encoding/json"
	"fmt"
	model "golang-mongodb-restful-starter-kit/app/models"
	"golang-mongodb-restful-starter-kit/app/services/auth"
	"golang-mongodb-restful-starter-kit/app/services/jwt"
	"golang-mongodb-restful-starter-kit/config"
	"golang-mongodb-restful-starter-kit/utility"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// AuthHandler
type AuthHadler struct {
	au auth.AuthService
	c  *config.Configuration
}

// AuthRouter
func AuthRouter(au auth.AuthService, c *config.Configuration, router *mux.Router) {

	authHandler := &AuthHadler{au, c}
	// ------------------------- Auth APIs ------------------------------
	router.HandleFunc(BaseRoute+"/auth/register", authHandler.Create).Methods(http.MethodPost)
	router.HandleFunc(BaseRoute+"/auth/login", authHandler.Login).Methods(http.MethodPost)

}

// Create
func (h *AuthHadler) Create(w http.ResponseWriter, r *http.Request) {
	requestUser := new(model.User)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)
	result := make(map[string]interface{})
	if validateError := requestUser.Validate(); validateError != nil {
		fmt.Println(validateError)
		result = utility.NewHTTPCustomError(utility.BadRequest, validateError.Error(), http.StatusBadRequest)
		utility.Response(w, result)
		return
	}

	requestUser.Initialize()

	if h.au.IsUserAlreadyExists(r.Context(), requestUser.Email) {
		result = utility.NewHTTPError(utility.UserAlreadyExists, http.StatusBadRequest)
		utility.Response(w, result)
		return
	}
	err := h.au.Create(r.Context(), requestUser)
	if err != nil {
		result = utility.NewHTTPError(utility.EntityCreationError, http.StatusBadRequest)
	} else {
		result["message"] = "Successfully Registered"
	}
	utility.Response(w, result)
}

func (h *AuthHadler) Login(w http.ResponseWriter, r *http.Request) {
	credentials := new(model.Credential)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&credentials)

	user, err := h.au.Login(r.Context(), credentials)
	result := make(map[string]interface{})
	if err != nil || user == nil {
		log.Println(err)
		result = utility.NewHTTPError(utility.Unauthorized, http.StatusBadRequest)
		utility.Response(w, result)
		return
	}
	j := jwt.JwtToken{C: h.c}
	tokenMap, err := j.CreateToken(user.ID.Hex(), user.Role)
	if err != nil {
		log.Println(err)
		result = utility.NewHTTPError(utility.InternalError, 501)
		utility.Response(w, result)
		return
	}

	result["token"] = tokenMap["token"]
	result["user"] = user
	utility.Response(w, result)
}
