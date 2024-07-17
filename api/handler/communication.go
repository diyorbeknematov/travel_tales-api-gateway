package handler

import (
	"api-gateway/generated/communication"
	"api-gateway/models"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// @Summary Send Message to User
// @Description Send a message to a user
// @Tags Messages
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param message body communication.SendMessageRequest true "Message"
// @Success 200 {object} communication.SendMessageResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/messages [post]
func (h *Handler) SendMessageUser(ctx *gin.Context) {
	var req communication.SendMessageRequest
	id, exists := ctx.Get("user_id")
	if !exists {
		h.Logger.Error("Error in get user id")
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Message: "Error in get user id",
		})
		return
	}

	req.SendeId = cast.ToString(id)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Error bind json", slog.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Message: "Error bind json",
		})
		return
	}

	resp, err := h.CommunityClient.SendMessageUser(ctx, &req)
	if err != nil {
		h.Logger.Error("xatolik xabarni yuborishda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik xabarni yuborishda",
		})
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

// @Summary List Messages
// @Description List all messages
// @Tags Messages
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Page limit"
// @Success 200 {object} communication.ListMessageResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/messages [get]
func (h *Handler) ListMessage(ctx *gin.Context) {
	var req communication.ListMessageRequest
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

	resp, err := h.CommunityClient.ListMessage(ctx, &req)
	if err != nil {
		h.Logger.Error("xatolik xabarlarni ro'yxatlashda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik xabarlarni ro'yxatlashda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Add Travel Tips
// @Description Add new travel tips
// @Tags TravelTips
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param tips body communication.AddTravelTipsRequest true "Travel Tips"
// @Success 200 {object} communication.AddTravelTipsResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/travel-tips [post]
func (h *Handler) AddTravelTips(ctx *gin.Context) {
	var req communication.AddTravelTipsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Error bind json", slog.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Message: "Error bind json",
		})
		return
	}

	resp, err := h.CommunityClient.AddTravelTips(ctx, &req)
	if err != nil {
		h.Logger.Error("xatolik sayohat tavsiyalarini qo'shishda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik sayohat tavsiyalarini qo'shishda",
		})
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

// @Summary Get Travel Tips
// @Description Get travel tips by ID
// @Tags TravelTips
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Page limit"
// @Param category query string true "Category"
// @Success 200 {object} communication.GetTravelTipsResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/travel-tips/ [get]
func (h *Handler) GetTravelTips(ctx *gin.Context) {
	var req communication.GetTravelTipsRequest
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

	resp, err := h.CommunityClient.GetTravelTips(ctx, &req)
	if err != nil {
		h.Logger.Error("xatolik sayohat tavsiyalarini olishda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik sayohat tavsiyalarini olishda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get User Statistics
// @Description Get user statistics
// @Tags User
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} communication.GetUserStaticsResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/users/{userId}/statics [get]
func (h *Handler) GetUserStatics(ctx *gin.Context) {
	userId := ctx.Param("userId")
	req := &communication.GetUserStaticsRequest{UserId: userId}

	resp, err := h.CommunityClient.GetUserStatics(ctx, req)
	if err != nil {
		h.Logger.Error("xatolik foydalanuvchi statistikasini olishda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik foydalanuvchi statistikasini olishda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
