package routes

import (
	api "golang-mongodb-restful-starter-kit/app/handlers"
	userRepo "golang-mongodb-restful-starter-kit/app/repositories/user"
	authSrv "golang-mongodb-restful-starter-kit/app/services/auth"
	userSrv "golang-mongodb-restful-starter-kit/app/services/user"
	"golang-mongodb-restful-starter-kit/config"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

var (
	BaseRoute = "/api/v1"
)

func InitializeRoutes(router *mux.Router, dbSession *mgo.Session, conf *config.Configuration) {
	userRepository := userRepo.New(dbSession, conf)
	userService := userSrv.New(userRepository)
	authService := authSrv.New(userRepository)
	authAPI := api.NewAuthAPI(authService, conf)
	userAPI := api.NewUserAPI(userService)

	// Routes

	//  -------------------------- Auth APIs ------------------------------------
	router.HandleFunc(BaseRoute+"/auth/register", authAPI.Create).Methods(http.MethodPost)
	router.HandleFunc(BaseRoute+"/auth/login", authAPI.Login).Methods(http.MethodPost)

	// -------------------------- User APIs ------------------------------------
	router.HandleFunc(BaseRoute+"/users/me", userAPI.Get).Methods(http.MethodGet)
	router.HandleFunc(BaseRoute+"/users", userAPI.Update).Methods(http.MethodPut)

}
