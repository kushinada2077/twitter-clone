package router

import (
	"net/http"
	"twitter-clone/handlers/auth"
	"twitter-clone/handlers/follow"
	"twitter-clone/repositories"
	"twitter-clone/services"

	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *http.ServeMux {
	mux := http.NewServeMux()

	userRepo := repositories.NewUserRepository(db)
	followRepo := repositories.NewFollowRepository(db)

	followService := services.NewFollowService(followRepo, userRepo)
	followHandler := follow.NewFollowHandler(followService)

	authService := services.NewAuthService(userRepo)
	authHandler := auth.NewAuthHandler(authService)

	RegisterAuthRoutes(mux, authHandler)
	RegisterFollowRoutes(mux, followHandler)
	return mux
}
