package router

import (
	"net/http"
	"twitter-clone/handlers/auth"
)

func RegisterAuthRoutes(mux *http.ServeMux, authHandler *auth.AuthHandler) {
	mux.HandleFunc("/auth/signup", authHandler.SignupHandler)
	mux.HandleFunc("/auth/login", authHandler.LoginHandler)
}
