package main

import (
	"golang-mongodb-restful-starter-kit/api"
	"golang-mongodb-restful-starter-kit/config"
	"golang-mongodb-restful-starter-kit/service/auth"
	"golang-mongodb-restful-starter-kit/service/jwt"
	"golang-mongodb-restful-starter-kit/service/user"
	"golang-mongodb-restful-starter-kit/store"
	"golang-mongodb-restful-starter-kit/utility"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

func main() {

	// Initialize config
	c := config.NewConfig()

	// Make connection with db and get instance
	db := store.GetInstance(c)
	db.SetSafe(&mgo.Safe{})
	userService := user.New(db, c)
	authService := auth.New(db, c)

	// Router
	router := mux.NewRouter()

	api.UserRouter(userService, router)
	api.AuthRouter(authService, c, router)

	//
	jwtService := jwt.JwtToken{C: c}

	// Added middleware over all request to authenticate
	router.Use(jwtService.ProtectedEndpoint)

	// Server configuration
	srv := &http.Server{
		Handler:      utility.Headers(router), // Set header to routes
		Addr:         c.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Application is running at ", c.Address)

	// Serving application at specified port
	log.Fatal(srv.ListenAndServe())

}
