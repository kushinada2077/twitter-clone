package router

import (
	"net/http"
	"twitter-clone/handlers"
)

func RegisterFollowRoutes(mux *http.ServeMux, followHandler *handlers.FollowHandler) {
	mux.HandleFunc("/users/{followeeID}/follow", followHandler.Follow)
	mux.HandleFunc("/users/{followeeID}/unfollow", followHandler.Unfollow)
}
