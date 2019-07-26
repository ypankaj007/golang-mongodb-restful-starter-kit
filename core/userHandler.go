package core

import (
	"go-restapis/core/httphandler"
	"go-restapis/service/user"
	"go-restapis/utility"
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
		httphandler.Response(w, httphandler.NewHttpError(httphandler.InternalError, 500))
	} else {
		httphandler.Response(w, user)
	}
}

func (h *UserHadler) Update(w http.ResponseWriter, r *http.Request) {

}
