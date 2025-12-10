package domain

import (
	"net/http"
)

type emptyUsernameOrPassWordErr struct{}

func (e emptyUsernameOrPassWordErr) Error() string { return "username and password required" }
func (e emptyUsernameOrPassWordErr) Status() int   { return http.StatusBadRequest }

type shortPasswordLengthErr struct{}

func (e shortPasswordLengthErr) Error() string { return "password must be at least 6 characters" }
func (e shortPasswordLengthErr) Status() int   { return http.StatusBadRequest }

type usernameAlreadyExistsErr struct{}

func (e usernameAlreadyExistsErr) Error() string { return "username already exists" }
func (e usernameAlreadyExistsErr) Status() int   { return http.StatusConflict }

var (
	ErrEmptyUsernameOrPassword = emptyUsernameOrPassWordErr{}
	ErrShortPasswordLength     = shortPasswordLengthErr{}
	ErrUsernameAlredayExists   = usernameAlreadyExistsErr{}
)
