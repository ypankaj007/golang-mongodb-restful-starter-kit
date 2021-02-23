package handlers

import (
	"encoding/json"
	"golang-mongodb-restful-starter-kit/app/models"
	"golang-mongodb-restful-starter-kit/app/services/auth"
	"golang-mongodb-restful-starter-kit/app/services/jwt"
	"golang-mongodb-restful-starter-kit/config"
	"golang-mongodb-restful-starter-kit/utility"
	"log"
	"net/http"
)

// AuthHandler ..
type AuthHandler struct {
	au auth.AuthServiceInterface
	c  *config.Configuration
}

func NewAuthAPI(authSrv auth.AuthServiceInterface, conf *config.Configuration) *AuthHandler {
	return &AuthHandler{
		au: authSrv,
		c:  conf,
	}
}

// Create godoc
// @Summary Register user
// @Description Register user api if not exists
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   payload     body    signupReq     true        "User Data"
// @Success 200 {object} basicResponse
// @Success 200 {object} errorRes
// @Router /auth/register [post]
func (h *AuthHandler) Create(w http.ResponseWriter, r *http.Request) {
	payload := new(signupReq)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&payload)
	requestUser := &models.User{Name: payload.Name, Email: payload.Email, Password: payload.Password}
	result := make(map[string]interface{})
	if validateError := requestUser.Validate(); validateError != nil {
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

		result = utility.SuccessPayload(nil, "Successfully registered", 201)
	}
	utility.Response(w, result)
}

// Login godoc
// @Summary Login user
// @Description Login user api with email and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   payload     body    models.Credential     true        "User Data"
// @Success 200 {object} loginRes
// @Success 200 {object} errorRes
// @Router /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	credentials := new(models.Credential)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&credentials)

	user, err := h.au.Login(r.Context(), credentials)
	if err != nil || user == nil {
		log.Println(err)
		result := utility.NewHTTPError(utility.Unauthorized, http.StatusBadRequest)
		utility.Response(w, result)
		return
	}
	j := jwt.JwtToken{C: h.c}
	tokenMap, err := j.CreateToken(user.ID.Hex(), user.Role)
	if err != nil {
		log.Println(err)
		result := utility.NewHTTPError(utility.InternalError, 501)
		utility.Response(w, result)
		return
	}

	res := &loginRes{
		Token: tokenMap["token"],
		User:  user,
	}
	result := utility.SuccessPayload(res, "Successfully loggedIn")
	utility.Response(w, result)
}
