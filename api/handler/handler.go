package handler

import (
	"api-gateway/generated/communication"
	"api-gateway/generated/destination"
	"api-gateway/generated/itineraries"
	"api-gateway/generated/stories"
	"api-gateway/generated/user"
	"api-gateway/service"
	"log/slog"
)

type Handler struct {
	UserClient        user.AuthServiceClient
	StoriesClient     stories.TravelStoriesServiceClient
	ItineraryClient   itineraries.ItinerariesServiceClient
	CommunityClient   communication.CommunicationServiceClient
	DestinationClient destination.TravelDestinationServiceClient
	Logger            *slog.Logger
}

func NewHandler(serviceManager service.IServiceManager, logger *slog.Logger) *Handler {
	return &Handler{
		UserClient: serviceManager.UserService(),
		StoriesClient: serviceManager.StoriesService(),
		ItineraryClient: serviceManager.ItinerariesService(),
		CommunityClient: serviceManager.CommunicationService(),
		DestinationClient: serviceManager.TravelDestinationService(),
		Logger: logger,
	}
}