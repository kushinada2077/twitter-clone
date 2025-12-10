package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"twitter-clone/pkg/models"
	"twitter-clone/pkg/types"

	"twitter-clone/utils"

	"gorm.io/gorm"
)

func SignupHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.SignupRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.Error(w, http.StatusBadRequest, "invalid JSON", err)
			return
		}

		req.Username = strings.TrimSpace(req.Username)
		if req.Username == "" || req.Password == "" {
			utils.Error(w, http.StatusBadRequest, "username and password required")
			return
		}
		if len(req.Password) < 6 {
			utils.Error(w, http.StatusBadRequest, "password must be at least 6 characters")
			return
		}

		var count int64
		if err := db.Model(&models.User{}).Where("username = ?", req.Username).Count(&count).Error; err != nil {
			utils.Error(w, http.StatusInternalServerError, "server error", err)
			return
		}
		if count > 0 {
			utils.Error(w, http.StatusConflict, "username already exists")
			return
		}

		hashed, err := utils.HashPassWord(req.Password)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, "server error", err)
			return
		}

		user := models.User{
			Username:     req.Username,
			PasswordHash: hashed,
		}
		if err := db.Create(&user).Error; err != nil {
			if strings.Contains(err.Error(), "Duplicate") || strings.Contains(err.Error(), "duplicate") {
				utils.Error(w, http.StatusConflict, "Username already exists")
				return
			}
			utils.Error(w, http.StatusInternalServerError, "server error", err)
			return
		}

		utils.JSON(w, http.StatusCreated, types.SignupResponse{
			Message: "user created",
			UserID:  user.ID,
		})
	}
}

func LoginHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.Error(w, http.StatusBadRequest, "invalid JSON", err)
			return
		}

		req.Username = strings.TrimSpace(req.Username)
		if req.Username == "" || req.Password == "" {
			utils.Error(w, http.StatusBadRequest, "invalid username or password")
			return
		}

		user := models.User{
			Username: req.Username,
		}

		if err := db.Where("username = ?", user.Username).First(&user).Error; err != nil {
			utils.Error(w, http.StatusUnauthorized, "invalid username or password", err)
			return
		}

		if err := utils.ComparePassword(user.PasswordHash, req.Password); err != nil {
			utils.Error(w, http.StatusUnauthorized, "invalid username or password")
			return
		}

		token, err := utils.GenerateToken(user.ID)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, "server error", err)
			return
		}

		utils.JSON(w, http.StatusOK, types.LoginResponse{Token: token})
	}
}
