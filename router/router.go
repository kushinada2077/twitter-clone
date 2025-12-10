package router

import (
	"net/http"
	"twitter-clone/handlers"
	"twitter-clone/repositories"
	"twitter-clone/services"

	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *http.ServeMux {
	mux := http.NewServeMux()

	userRepo := repositories.NewUserRepository(db)
	followRepo := repositories.NewFollowRepository(db)

	followService := services.NewFollowService(followRepo, userRepo)
	followHandler := handlers.NewFollowHandler(followService)

	RegisterAuthRoutes(mux, db)
	RegisterFollowRoutes(mux, followHandler)
	return mux
}
