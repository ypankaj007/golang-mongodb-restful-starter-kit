package main

import (
	"go-restapis/config"
	"go-restapis/core"
	"go-restapis/core/httphandler"
	"log"
	"net/http"
	"time"

	"go-restapis/service/auth"
	"go-restapis/service/jwt"
	"go-restapis/service/user"
	"go-restapis/store"

	"github.com/gorilla/mux"
)

func main() {

	c := config.NewConfig()

	db := store.GetInstance(c)

	userService := user.New(db)
	authService := auth.New(db)

	userHandler := core.NewUserHandler(userService)
	authHandler := core.NewAuthHandler(authService, c)
	jwtService := jwt.JwtToken{C: c}

	// Router
	router := mux.NewRouter()

	// Auth APIs
	router.HandleFunc("/api/v1/auth/register", authHandler.Create).Methods("POST")
	router.HandleFunc("/api/v1/auth/login", authHandler.Login).Methods("POST")

	// User APIs
	router.HandleFunc("/api/v1/users/{userId}", userHandler.Get).Methods("GET")
	router.HandleFunc("/api/v1/users", userHandler.Update).Methods("PUT")

	// Middleware
	router.Use(jwtService.ProtectedEndpoint)

	// Server
	srv := &http.Server{
		Handler:      httphandler.Headers(router),
		Addr:         c.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Application is running at ", c.Address)
	log.Fatal(srv.ListenAndServe())

}
