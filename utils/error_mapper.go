package utils

import (
	"errors"
	"net/http"
	"strconv"
	"twitter-clone/pkg/domain"
)

func HandleAPIError(w http.ResponseWriter, err error) {
	var code int
	var message string

	if getter, ok := err.(domain.HTTPStatusGetter); ok {
		code = getter.Status()
		message = getter.Error()
	} else {
		switch {
		case errors.Is(err, strconv.ErrRange) || errors.Is(err, strconv.ErrSyntax):
			code = http.StatusBadRequest
			message = "invalid format or range for an ID"

		default:
			code = http.StatusInternalServerError
			message = "server error"
			Error(w, code, message, err)
			return
		}
	}

	Error(w, code, message)
}
