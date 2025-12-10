package types

type contextKey struct{}

var UserIDKey = contextKey{}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Message string `json:"message"`
	UserID  uint   `json:"user_id,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type FollowResponse struct {
	Message string `json:"message"`
}

type UnfollowResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
