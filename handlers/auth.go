package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"twitter-clone/models"
	"twitter-clone/utils"

	"gorm.io/gorm"
)

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Message string `json:"message"`
	UserID  uint   `json:"user_id,omitempty"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func SignupHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req SignupRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		req.Username = strings.TrimSpace(req.Username)
		if req.Username == "" || req.Password == "" {
			http.Error(w, "username and password required", http.StatusBadRequest)
			return
		}
		if len(req.Password) < 6 {
			http.Error(w, "password must be at least 6 characters", http.StatusBadRequest)
		}

		var count int64
		if err := db.Model(&models.User{}).Where("username = ?", req.Username).Count(&count).Error; err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		if count > 0 {
			http.Error(w, "username already exists", http.StatusConflict)
		}

		hashed, err := utils.HashPassWord(req.Password)
		if err != nil {
			http.Error(w, "failed to hash password", http.StatusInternalServerError)
			return
		}

		user := models.User{
			Username:     req.Username,
			PasswordHash: hashed,
		}
		if err := db.Create(&user).Error; err != nil {
			if strings.Contains(err.Error(), "Duplicate") || strings.Contains(err.Error(), "duplicate") {
				http.Error(w, "username already exists", http.StatusConflict)
				return
			}
			http.Error(w, "failed to create user", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(SignupResponse{
			Message: "user created",
			UserID:  user.ID,
		})
	}
}

func LoginHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		req.Username = strings.TrimSpace(req.Username)
		if req.Username == "" || req.Password == "" {
			http.Error(w, "invalid username or password", http.StatusBadRequest)
			return
		}

		user := models.User{
			Username: req.Username,
		}

		if err := db.Where("username = ?", user.Username).First(&user).Error; err != nil {
			http.Error(w, "invalid username or password", http.StatusUnauthorized)
			return
		}

		if err := utils.ComparePassword(user.PasswordHash, req.Password); err != nil {
			http.Error(w, "invalid username or password", http.StatusUnauthorized)
			return
		}

		token := "token_user_" + fmt.Sprint(user.ID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(LoginResponse{
			Token: token,
		})
	}
}
