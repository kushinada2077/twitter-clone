package domainerrors

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

type usernameNotExistsErr struct{}

func (e usernameNotExistsErr) Error() string { return "username doesn't exist" }
func (e usernameNotExistsErr) Status() int   { return http.StatusUnauthorized }

type invalidUsernameOrPasswordErr struct{}

func (e invalidUsernameOrPasswordErr) Error() string { return "invalid username or password" }
func (e invalidUsernameOrPasswordErr) Status() int   { return http.StatusUnauthorized }

var (
	ErrEmptyUsernameOrPassword   = emptyUsernameOrPassWordErr{}
	ErrShortPasswordLength       = shortPasswordLengthErr{}
	ErrUsernameAlredayExists     = usernameAlreadyExistsErr{}
	ErrUsernameNotExists         = usernameNotExistsErr{}
	ErrInvalidUsernameOrPassword = invalidUsernameOrPasswordErr{}
)
