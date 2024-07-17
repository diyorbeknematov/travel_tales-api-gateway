package handler

import (
	"api-gateway/generated/user"
	"api-gateway/models"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// @Summary Get user profile
// @Description Get user profile with user id
// @Tags User
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} user.GetProfileResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/users/profile [get]
func (h *Handler) GetUserProfileHandle(ctx *gin.Context) {
	id, exists := ctx.Get("user_id")
	if !exists {
		h.Logger.Error("Error in get user id")
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Message: "Error in get user id",
		})
		return
	}

	resp, err := h.UserClient.GetUserProfile(ctx, &user.GetProfileRequest{Id: cast.ToString(id)})
	if err != nil {
		h.Logger.Error("Error in get user profile", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "Error in get user profile",
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update user profile
// @Description Update user profile with user id
// @Tags User
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param UserProfile body user.UpdateProfileRequest true "UserProfile"
// @Success 200 {object} user.UpdateProfileResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/users/profile [put]
func (h *Handler) UpdateUserProfileHandle(ctx *gin.Context) {
	id, exists := ctx.Get("user_id")
	var profile user.UpdateProfileRequest
	if !exists {
		h.Logger.Error("Error in get user id")
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Message: "Error in get user id",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&profile); err != nil {
		h.Logger.Error("error bind json", slog.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Message: "Error in bind json",
		})
		return
	}

	profile.Id = cast.ToString(id)

	resp, err := h.UserClient.UpdateUserProfile(ctx, &profile)
	if err != nil {
		h.Logger.Error("Error in updated user", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "Error in updated user",
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary List users
// @Description GET List users
// @Tags User
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param page query int true "Page number"
// @Param limit query int true "Page limit"
// @Success 200 {object} user.ListUsersResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/users [get]
func (h *Handler) ListUsersHandle(ctx *gin.Context) {
	var req user.ListUsersRequest

	// Get 'page' query parameter and convert it to int32
	pageStr := ctx.Query("page")
	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			h.Logger.Error("Error converting page to int", slog.String("error", err.Error()))
			ctx.JSON(http.StatusBadRequest, models.Errors{
				Message: "Error parsing page from string to int",
			})
			return
		}
		req.Page = int32(page)
	} else {
		req.Page = 1
	}

	// Get 'limit' query parameter and convert it to int32
	limitStr := ctx.Query("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			h.Logger.Error("Error converting limit to int", slog.String("error", err.Error()))
			ctx.JSON(http.StatusBadRequest, models.Errors{
				Message: "Error parsing limit from string to int",
			})
			return
		}
		req.Limit = int32(limit)
	}

	// Call the ListUsers method on the gRPC client
	resp, err := h.UserClient.ListUsers(ctx, &req)
	if err != nil {
		h.Logger.Error("Error in get list users", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "Error in get list users",
		})
		return
	} else {
		req.Limit = 10
	}

	// Return the response
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete user
// @Description Delete user by ID
// @Tags User
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} user.DeleteUserResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/users/{userId} [delete]
func (h *Handler) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("userId")

	resp, err := h.UserClient.DeleteUser(ctx, &user.DeleteUserRequest{Id: userId})
	if err != nil {
		h.Logger.Error("Error deleting user", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "Error deleting user",
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get user activity
// @Description Get user activity by user ID
// @Tags User
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} user.GetUserActivityResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/users/{userId}/activity [get]
func (h *Handler) GetUserActivity(ctx *gin.Context) {
	userId := ctx.Param("userId")

	resp, err := h.UserClient.GetUserActivity(ctx, &user.GetUserActivityRequest{Id: userId})
	if err != nil {
		h.Logger.Error("Error getting user activity", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "Error getting user activity",
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Follow user
// @Description Follow a user by ID
// @Tags User
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} user.FollowUserResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/users/{userId}/follow [post]
func (h *Handler) FollowUser(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, exists := ctx.Get("user_id")
	if !exists {
		h.Logger.Error("Error in get user id")
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Message: "Error in get user id",
		})
		return
	}
	resp, err := h.UserClient.FollowUser(ctx, &user.FollowUserRequest{FollowerId: userId, FollowingId: cast.ToString(id)})
	if err != nil {
		h.Logger.Error("Error following user", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "Error following user",
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary List followers
// @Description List followers of a user by user ID
// @Tags User
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} user.ListFollowersResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/users/{userId}/followers [get]
func (h *Handler) ListFollowers(ctx *gin.Context) {
	req := user.ListFollowersRequest{}
	userId := ctx.Param("userId")

	// Get 'page' query parameter and convert it to int32
	pageStr := ctx.Query("page")
	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			h.Logger.Error("Error converting page to int", slog.String("error", err.Error()))
			ctx.JSON(http.StatusBadRequest, models.Errors{
				Message: "Error parsing page from string to int",
			})
			return
		}
		req.Page = int32(page)
	} else {
		req.Page = 1
	}

	// Get 'limit' query parameter and convert it to int32
	limitStr := ctx.Query("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			h.Logger.Error("Error converting limit to int", slog.String("error", err.Error()))
			ctx.JSON(http.StatusBadRequest, models.Errors{
				Message: "Error parsing limit from string to int",
			})
			return
		}
		req.Limit = int32(limit)
	} else {
		req.Limit = 10
	}

	resp, err := h.UserClient.ListFollowers(ctx, &user.ListFollowersRequest{UserId: userId})
	if err != nil {
		h.Logger.Error("Error listing followers", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "Error listing followers",
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
