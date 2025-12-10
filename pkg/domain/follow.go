package domain

import (
	"net/http"
)

type userNotFoundErr struct{}

func (e userNotFoundErr) Error() string { return "user not found" }
func (e userNotFoundErr) Status() int   { return http.StatusNotFound }

type followNotFoundErr struct{}

func (e followNotFoundErr) Error() string { return "follow not found" }
func (e followNotFoundErr) Status() int   { return http.StatusNotFound }

type alreadyFollowingErr struct{}

func (e alreadyFollowingErr) Error() string { return "already following" }
func (e alreadyFollowingErr) Status() int   { return http.StatusConflict }

type cannotFollowYourselfErr struct{}

func (e cannotFollowYourselfErr) Error() string { return "cannot follow yourself" }
func (e cannotFollowYourselfErr) Status() int   { return http.StatusBadRequest }

type cannotUnfollowYourselfErr struct{}

func (e cannotUnfollowYourselfErr) Error() string { return "cannot unfollow yourself" }
func (e cannotUnfollowYourselfErr) Status() int   { return http.StatusBadRequest }

var (
	ErrUserNotFound           = userNotFoundErr{}
	ErrFollowNotFound         = followNotFoundErr{}
	ErrAlreadyFollowing       = alreadyFollowingErr{}
	ErrCannotFollowYourself   = cannotFollowYourselfErr{}
	ErrCannotUnfollowYourself = cannotUnfollowYourselfErr{}
)
