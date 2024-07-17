package handler

import (
	"api-gateway/generated/destination"
	"api-gateway/models"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary List Travel Destinations
// @Description List all Travel Destinations
// @Tags Destinations
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param page query int fale "Page number"
// @Param limit query int false "Page limit"
// @Success 200 {object} destination.ListDetinationResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/destinations [get]
func (h *Handler) ListTravelDestinations(ctx *gin.Context) {
	var req destination.ListDetinationRequest
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

	resp, err := h.DestinationClient.ListTravelDestnations(ctx, &req)
	if err != nil {
		h.Logger.Error("xatolik sayohat manzillarini ro'yxatlashda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik sayohat manzillarini ro'yxatlashda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get Travel Destination
// @Description Get a single Travel Destination by ID
// @Tags Destinations
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param destinationId path string true "Destination ID"
// @Success 200 {object} destination.GetDestinationResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/destinations/{destinationId} [get]
func (h *Handler) GetTravelDestination(ctx *gin.Context) {
	destinationId := ctx.Param("destinationId")
	req := &destination.GetDestinationRequest{Id: destinationId}

	resp, err := h.DestinationClient.GetTravelDestination(ctx, req)
	if err != nil {
		h.Logger.Error("xatolik sayohat manzilini olishda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik sayohat manzilini olishda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get Trend Travel Destinations
// @Description Get trending Travel Destinations
// @Tags Destinations
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param Top query int false "Top Destination"
// @Success 200 {object} destination.GetTrendDestinationResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/destinations/trending [get]
func (h *Handler) GetTrendDestinations(ctx *gin.Context) {
	req := &destination.GetTrendDestinationRequest{}

	topStr := ctx.Query("Top")

	if topStr != "" {
		page, err := strconv.Atoi(topStr)
		if err != nil {
			h.Logger.Error("Error converting page to int", slog.String("error", err.Error()))
			ctx.JSON(http.StatusBadRequest, models.Errors{
				Message: "Error parsing page from string to int",
			})
			return
		}
		req.Limit = int32(page)
	} else {
		req.Limit = 10
	}

	resp, err := h.DestinationClient.GetTrendDestinations(ctx, req)
	if err != nil {
		h.Logger.Error("xatolik trend sayohat manzillarini olishda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik trend sayohat manzillarini olishda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
