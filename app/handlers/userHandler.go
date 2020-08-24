package handlers

import (
	"encoding/json"
	model "golang-mongodb-restful-starter-kit/app/models"
	"golang-mongodb-restful-starter-kit/app/services/user"
	"golang-mongodb-restful-starter-kit/utility"
	"net/http"

	"github.com/gorilla/mux"
)

// UserHandler - handles user request
type UserHandler struct {
	us user.UserService
}

// UserRouter
func UserRouter(us user.UserService, router *mux.Router) {
	userHandler := &UserHandler{us}

	// -------------------------- User APIs ------------------------------------
	router.HandleFunc(BaseRoute+"/users/me", userHandler.Get).Methods(http.MethodGet)
	router.HandleFunc(BaseRoute+"/users", userHandler.Update).Methods(http.MethodPut)

}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	user, err := h.us.Get(r.Context(), utility.GetLoggedInUserID(r))

	if err != nil {
		utility.Response(w, utility.NewHTTPError(utility.InternalError, 500))
	} else {
		utility.Response(w, user)
	}
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	updaateUser := new(model.UserUpdate)
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

	result["message"] = "Successfully updated"
	utility.Response(w, result)

}
