package core

import (
	"golang-mongodb-restful-starter-kit/core/httphandler"
	"golang-mongodb-restful-starter-kit/service/user"
	"golang-mongodb-restful-starter-kit/utility"
	"net/http"
)

// UserHandler
type UserHadler struct {
	us user.UserService
}

// NewUserHandler
func NewUserHandler(us user.UserService) *UserHadler {
	return &UserHadler{us}

}

func (h *UserHadler) Get(w http.ResponseWriter, r *http.Request) {
	user, err := h.us.Get(r.Context(), utility.GetLoggedInUserId(r))

	if err != nil {
		httphandler.Response(w, httphandler.NewHTTPError(httphandler.InternalError, 500))
	} else {
		httphandler.Response(w, user)
	}
}

func (h *UserHadler) Update(w http.ResponseWriter, r *http.Request) {

}
