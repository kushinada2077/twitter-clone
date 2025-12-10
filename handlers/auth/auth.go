package auth

import (
	"net/http"
	"strings"
	"twitter-clone/services"

	"twitter-clone/utils"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(s services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: s,
	}
}

func (h *AuthHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid JSON", err)
		return
	}

	req.Username = strings.TrimSpace(req.Username)

	userID, err := h.authService.Signup(req.Username, req.Password)
	if err != nil {
		utils.HandleAPIError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, SignupResponse{
		Message: "user created",
		UserID:  userID,
	})
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid JSON", err)
		return
	}

	req.Username = strings.TrimSpace(req.Username)

	token, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		utils.HandleAPIError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, LoginResponse{Token: token})
}
