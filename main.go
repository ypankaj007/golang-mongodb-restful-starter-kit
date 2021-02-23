package main

import (
	"golang-mongodb-restful-starter-kit/app/middleware"
	"golang-mongodb-restful-starter-kit/app/services/jwt"
	"golang-mongodb-restful-starter-kit/config"
	"golang-mongodb-restful-starter-kit/db"
	_ "golang-mongodb-restful-starter-kit/docs"
	"golang-mongodb-restful-starter-kit/routes"
	"golang-mongodb-restful-starter-kit/utility"

	httpSwagger "github.com/swaggo/http-swagger"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

// @title Application API
// @version 1.0
// @description Auth apis (signup/login) and user apis
// @contact.name API Support
// @contact.email ypankaj007@gmail.com
// @license.name Apache 2.0
// @host localhost:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/v1
func main() {

	// Initialize config
	conf := config.NewConfig()

	// Make connection with db and get instance
	dbSession := db.GetInstance(conf)

	//
	dbSession.SetSafe(&mgo.Safe{})

	// Router
	router := mux.NewRouter()
	routes.InitializeRoutes(router, dbSession, conf)
	// JWT services
	jwtService := jwt.JwtToken{C: conf}

	// Added middleware over all request to authenticate
	router.Use(middleware.Cors, jwtService.ProtectedEndpoint)

	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	// Server configuration
	srv := &http.Server{
		Handler:      utility.Headers(router), // Set header to routes
		Addr:         conf.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Application is running at ", conf.Address)

	// Serving application at specified port
	log.Fatal(srv.ListenAndServe())

}
