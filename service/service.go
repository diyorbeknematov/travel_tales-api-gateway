package service

import (
	"api-gateway/config"
	"api-gateway/generated/communication"
	"api-gateway/generated/destination"
	"api-gateway/generated/itineraries"
	"api-gateway/generated/stories"
	"api-gateway/generated/user"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serviceManager struct {
	userClient        user.AuthServiceClient
	storiesClient     stories.TravelStoriesServiceClient
	itineraryClient   itineraries.ItinerariesServiceClient
	communityClient   communication.CommunicationServiceClient
	destinationClient destination.TravelDestinationServiceClient
}

type IServiceManager interface {
	UserService() user.AuthServiceClient
	StoriesService() stories.TravelStoriesServiceClient
	ItinerariesService() itineraries.ItinerariesServiceClient
	CommunicationService() communication.CommunicationServiceClient
	TravelDestinationService() destination.TravelDestinationServiceClient
}

func (s *serviceManager) UserService() user.AuthServiceClient {
	return s.userClient
}

func (s *serviceManager) StoriesService() stories.TravelStoriesServiceClient {
	return s.storiesClient
}

func (s *serviceManager) ItinerariesService() itineraries.ItinerariesServiceClient {
	return s.itineraryClient
}

func (s *serviceManager) CommunicationService() communication.CommunicationServiceClient {
	return s.communityClient
}

func (s *serviceManager) TravelDestinationService() destination.TravelDestinationServiceClient {
	return s.destinationClient
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {

	connUser, err := grpc.NewClient(
		fmt.Sprintf("travel-auth:%s", conf.USER_SERVICE_PORT),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connContent, err := grpc.NewClient(
		fmt.Sprintf("travel-content:%s", conf.CONTENT_SERVICE_PORT),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}


	serviceManager := &serviceManager{
		userClient:        user.NewAuthServiceClient(connUser),
		storiesClient:     stories.NewTravelStoriesServiceClient(connContent),
		itineraryClient:   itineraries.NewItinerariesServiceClient(connContent),
		communityClient:   communication.NewCommunicationServiceClient(connContent),
		destinationClient: destination.NewTravelDestinationServiceClient(connContent),
	}

	return serviceManager, nil
}
