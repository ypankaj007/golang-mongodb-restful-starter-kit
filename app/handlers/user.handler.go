package handlers

import (
	"encoding/json"
	"golang-mongodb-restful-starter-kit/app/models"
	userSrv "golang-mongodb-restful-starter-kit/app/services/user"
	"golang-mongodb-restful-starter-kit/utility"
	"net/http"
)

// UserHandler - handles user request
type UserHandler struct {
	us userSrv.UserServiceInterface
}

func NewUserAPI(userService userSrv.UserServiceInterface) *UserHandler {
	return &UserHandler{
		us: userService,
	}
}

// Get godoc
// @Summary Get Profile
// @Description Get user profile info
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Success 200 {object} errorRes
// @Security ApiKeyAuth
// @Router /users/me [get]
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	user, err := h.us.Get(r.Context(), utility.GetLoggedInUserID(r))

	if err != nil {
		utility.Response(w, utility.NewHTTPError(utility.InternalError, 500))
	} else {
		utility.Response(w, utility.SuccessPayload(user, ""))
	}
}

// Update godoc
// @Summary Update user
// @Description Update user info
// @Tags users
// @Accept  json
// @Produce  json
// @Param   payload     body    models.UserUpdate     true        "User Data"
// @Success 200 {object} basicResponse
// @Success 200 {object} errorRes
// @Security ApiKeyAuth
// @Router /users [put]
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	updaateUser := new(models.UserUpdate)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&updaateUser)
	result := make(map[string]interface{})
	err := h.us.Update(r.Context(), utility.GetLoggedInUserID(r), updaateUser)
	if err != nil {
		result = utility.NewHTTPCustomError(utility.BadRequest, err.Error(), http.StatusBadRequest)
		utility.Response(w, result)
		return
	}

	result = utility.SuccessPayload(nil, "Successfully updated")
	utility.Response(w, result)

}
