package auth

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
