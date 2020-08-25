package main

import (
	api "golang-mongodb-restful-starter-kit/app/handlers"
	"golang-mongodb-restful-starter-kit/app/middleware"
	"golang-mongodb-restful-starter-kit/app/services/auth"
	"golang-mongodb-restful-starter-kit/app/services/jwt"
	"golang-mongodb-restful-starter-kit/app/services/user"
	"golang-mongodb-restful-starter-kit/config"
	"golang-mongodb-restful-starter-kit/db"
	_ "golang-mongodb-restful-starter-kit/docs"
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
	c := config.NewConfig()

	// Make connection with db and get instance
	dbSession := db.GetInstance(c)

	//
	dbSession.SetSafe(&mgo.Safe{})

	// UserService
	userService := user.New(dbSession, c)

	// AuthService
	authService := auth.New(dbSession, c)

	// Router
	router := mux.NewRouter()

	// UserRouter
	api.UserRouter(userService, router)

	// AuthRouter
	api.AuthRouter(authService, c, router)

	// JWT services
	jwtService := jwt.JwtToken{C: c}

	// Added middleware over all request to authenticate
	router.Use(middleware.Cors, jwtService.ProtectedEndpoint)

	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

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

	// // Run our server in a goroutine so that it doesn't block.
	// go func() {
	//     if err := srv.ListenAndServe(); err != nil {
	//         log.Println(err)
	//     }
	// }()

	// c := make(chan os.Signal, 1)
	// // We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	// signal.Notify(c, os.Interrupt)

	// // Block until we receive our signal.
	// <-c

	// // Create a deadline to wait for.
	// ctx, cancel := context.WithTimeout(context.Background(), wait)
	// defer cancel()
	// // Doesn't block if no connections, but will otherwise wait
	// // until the timeout deadline.
	// srv.Shutdown(ctx)
	// // Optionally, you could run srv.Shutdown in a goroutine and block on
	// // <-ctx.Done() if your application should wait for other services
	// // to finalize based on context cancellation.
	// log.Println("shutting down")
	// os.Exit(0)

}
