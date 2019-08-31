package api

import (
	"golang-mongodb-restful-starter-kit/service/user"
	"golang-mongodb-restful-starter-kit/utility"
	"net/http"

	"github.com/gorilla/mux"
)

// UserHandler
type UserHadler struct {
	us user.UserService
}

// UserRouter
func UserRouter(us user.UserService, router *mux.Router) {
	userHandler := &UserHadler{us}

	// -------------------------- User APIs ------------------------------------
	router.HandleFunc(BaseRoute+"/users/{userId}", userHandler.Get).Methods(http.MethodGet)
	router.HandleFunc(BaseRoute+"/users", userHandler.Update).Methods(http.MethodPut)

}

func (h *UserHadler) Get(w http.ResponseWriter, r *http.Request) {
	user, err := h.us.Get(r.Context(), utility.GetLoggedInUserId(r))

	if err != nil {
		utility.Response(w, utility.NewHTTPError(utility.InternalError, 500))
	} else {
		utility.Response(w, user)
	}
}

func (h *UserHadler) Update(w http.ResponseWriter, r *http.Request) {

}
