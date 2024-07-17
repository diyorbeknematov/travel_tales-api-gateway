package handler

import (
	"api-gateway/generated/itineraries"
	"api-gateway/models"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Itinerary
// @Description Create new Itinerary
// @Tags Itineraries
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param Itinerary body itineraries.CreateItineraryRequest true "Itinerary"
// @Success 200 {object} itineraries.CreateItineraryResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/itineraries [post]
func (h *Handler) CreateItinerary(ctx *gin.Context) {
	var req itineraries.CreateItineraryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Error bind json", slog.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Message: "Error bind json",
		})
		return
	}

	resp, err := h.ItineraryClient.CreateItinerary(ctx, &req)
	if err != nil {
		h.Logger.Error("xatolik marshrutni yaratishda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik marshrutni yaratishda",
		})
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

// @Summary Update Itinerary
// @Description Update an existing Itinerary
// @Tags Itineraries
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param Itinerary body itineraries.UpdateItineraryRequest true "Itinerary"
// @Success 200 {object} itineraries.UpdateItineraryResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/itineraries/{itineraryId} [put]
func (h *Handler) UpdateItinerary(ctx *gin.Context) {
	var req itineraries.UpdateItineraryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Error bind json", slog.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Message: "Error bind json",
		})
		return
	}

	resp, err := h.ItineraryClient.UpdateItinerary(ctx, &req)
	if err != nil {
		h.Logger.Error("xatolik marshrutni yangilashda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik marshrutni yangilashda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete Itinerary
// @Description Delete an existing Itinerary
// @Tags Itineraries
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param itineraryId path string true "Itinerary ID"
// @Success 200 {object} itineraries.DeleteItineraryResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/itineraries/{itineraryId} [delete]
func (h *Handler) DeleteItinerary(ctx *gin.Context) {
	itineraryId := ctx.Param("itineraryId")
	req := &itineraries.DeleteItineraryRequest{Id: itineraryId}

	resp, err := h.ItineraryClient.DeleteItinerary(ctx, req)
	if err != nil {
		h.Logger.Error("xatolik marshrutni o'chirishda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik marshrutni o'chirishda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary List Itineraries
// @Description List all Itineraries
// @Tags Itineraries
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param page query int true "Page number"
// @Param limit query int true "Page limit"
// @Success 200 {object} itineraries.ListItinerariesResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/itineraries [get]
func (h *Handler) ListItineraries(ctx *gin.Context) {
	var req itineraries.ListItinerariesRequest
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

	resp, err := h.ItineraryClient.ListItineraries(ctx, &req)
	if err != nil {
		h.Logger.Error("xatolik marshrutlarni ro'yxatlashda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik marshrutlarni ro'yxatlashda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get Itinerary
// @Description Get a single Itinerary by ID
// @Tags Itineraries
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param itineraryId path string true "Itinerary ID"
// @Success 200 {object} itineraries.GetItineraryResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/itineraries/{itineraryId} [get]
func (h *Handler) GetItinerary(ctx *gin.Context) {
	itineraryId := ctx.Param("itineraryId")
	req := &itineraries.GetItineraryRequest{Id: itineraryId}

	resp, err := h.ItineraryClient.GetItinerary(ctx, req)
	if err != nil {
		h.Logger.Error("xatolik marshrutni olishda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik marshrutni olishda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary Leave Comment on Itinerary
// @Description Leave a comment on an Itinerary
// @Tags Itineraries
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param comment body itineraries.LeaveCommentRequest true "Comment"
// @Success 200 {object} itineraries.LeaveCommentResponse
// @Failure 400 {object} models.Errors
// @Failure 500 {object} models.Errors
// @Router /api/v1/itineraries/{itineraryId}/comments [post]
func (h *Handler) LeaveComment(ctx *gin.Context) {
	var req itineraries.LeaveCommentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Error bind json", slog.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Message: "Error bind json",
		})
		return
	}

	resp, err := h.ItineraryClient.LeaveComment(ctx, &req)
	if err != nil {
		h.Logger.Error("xatolik marshrutga izoh qo'shishda", slog.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Message: "xatolik marshrutga izoh qo'shishda",
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
