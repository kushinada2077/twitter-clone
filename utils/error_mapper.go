package utils

import (
	"errors"
	"net/http"
	"strconv"
	"twitter-clone/services"
)

func HandleAPIError(w http.ResponseWriter, err error) {
	var code int
	var message string

	switch {
	case errors.Is(err, services.ErrUserNotFound) || errors.Is(err, services.ErrFollowNotFound):
		code = http.StatusNotFound
		message = "The requested resource could not be found"

	case errors.Is(err, services.ErrCannotFollowYourself) || errors.Is(err, services.ErrCannotUnfollowYourself):
		code = http.StatusBadRequest
		message = "you cannot follow or unfollow yourself"

	case errors.Is(err, services.ErrAlreadyFollowing):
		code = http.StatusConflict
		message = "already following"

	case errors.Is(err, strconv.ErrRange) || errors.Is(err, strconv.ErrSyntax):
		code = http.StatusBadRequest
		message = "invalid format or range for an ID in the request path"

	default:
		code = http.StatusInternalServerError
		message = "server error"
	}

	Error(w, code, message)
}
