package router

import (
	"net/http"
	"twitter-clone/handlers"

	"gorm.io/gorm"
)

func RegisterAuthRoutesmux(mux *http.ServeMux, db *gorm.DB) {
	mux.HandleFunc("/auth/signup", handlers.SignupHandler(db))
	mux.HandleFunc("/auth/login", handlers.LoginHandler(db))
}
