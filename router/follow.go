package router

import (
	"net/http"
	"twitter-clone/handlers/follow"
)

func RegisterFollowRoutes(mux *http.ServeMux, followHandler *follow.FollowHandler) {
	mux.HandleFunc("/users/{followeeID}/follow", followHandler.Follow)
	mux.HandleFunc("/users/{followeeID}/unfollow", followHandler.Unfollow)
}
