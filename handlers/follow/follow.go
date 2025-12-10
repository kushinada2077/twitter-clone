package follow

import (
	"net/http"
	"twitter-clone/pkg/domain"
	"twitter-clone/services"
	"twitter-clone/utils"
)

type FollowHandler struct {
	followService services.FollowService
}

func NewFollowHandler(s services.FollowService) *FollowHandler {
	return &FollowHandler{
		followService: s,
	}
}

func (h *FollowHandler) Follow(w http.ResponseWriter, r *http.Request) {
	followerID, ok := r.Context().Value(domain.UserIDKey).(uint)
	if !ok {
		utils.Error(w, http.StatusUnauthorized, "invalid or missing token")
		return
	}

	followeeIDStr := r.PathValue("followeeID")
	followeeID, err := utils.ParseID(followeeIDStr)
	if err != nil {
		utils.HandleAPIError(w, err)
		return
	}

	if err := h.followService.Follow(followerID, followeeID); err != nil {
		utils.HandleAPIError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, FollowResponse{Message: "follow success"})
}

func (h *FollowHandler) Unfollow(w http.ResponseWriter, r *http.Request) {
	followerID, ok := r.Context().Value(domain.UserIDKey).(uint)
	if !ok {
		utils.Error(w, http.StatusUnauthorized, "invalid or missing token")
		return
	}

	followeeIDStr := r.PathValue("followeeID")
	followeeID, err := utils.ParseID(followeeIDStr)
	if err != nil {
		utils.HandleAPIError(w, err)
		return
	}

	if err := h.followService.Unfollow(followerID, followeeID); err != nil {
		utils.HandleAPIError(w, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, UnfollowResponse{Message: "unfollow success"})
}
