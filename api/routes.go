package api

import (
	"api-gateway/api/handler"
	"api-gateway/api/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "api-gateway/api/handler/docs"
)

// @title Auth Service API
// @version 1.0
// @description This is a sample server for Auth Service.
// @host localhost:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/v1
// @schemes http
func NewRouter(handle *handler.Handler) *gin.Engine {
	router := gin.Default()

	// Swagger endpointini sozlash
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(middleware.AuthMiddleware())
	router.Use(middleware.LoggerMiddleware())
	stories := router.Group("/api/v1/stories")
	{
		stories.POST("/", handle.CreateTravelStory)
		stories.PUT("/:storyId", handle.UpdateTravelStory)
		stories.DELETE("/:storyId", handle.DeleteTravelStory)
		stories.GET("/", handle.ListTravelStory)
		stories.GET("/:storyId", handle.GetTravelStory)
		stories.POST("/:storyId/comments", handle.AddComment)
		stories.GET("/:storyId/comments", handle.ListComments)
		stories.POST("/:storyId/like", handle.AddLike)
	}

	itinerary := router.Group("/api/v1/itineraries")
	{
		itinerary.POST("/", handle.CreateItinerary)
		itinerary.PUT("/:itineraryId", handle.UpdateItinerary)
		itinerary.DELETE("/:itineraryId", handle.DeleteItinerary)
		itinerary.GET("/", handle.ListItineraries)
		itinerary.GET("/:itineraryId", handle.GetItinerary)
		itinerary.POST("/:itineraryId/comments", handle.LeaveComment)
	}

	destinations := router.Group("/api/v1/destinations")
	{
		destinations.GET("/", handle.ListTravelDestinations)
		destinations.GET("/:destinationId", handle.GetTravelDestination)
		destinations.GET("/trending", handle.GetTrendDestinations)
	}

	message := router.Group("/api/v1/messages")
	{
		message.POST("/", handle.SendMessageUser)
		message.GET("/", handle.ListMessage)
	}

	tips := router.Group("/api/v1/travel-tips")
	{
		tips.POST("/", handle.AddTravelTips)
		tips.GET("/", handle.GetTravelTips)
	}

	user := router.Group("/api/v1/users")
	{
		user.GET("/profile", handle.GetUserProfileHandle)
		user.PUT("/profile", handle.UpdateUserProfileHandle)
		user.GET("/", handle.ListUsersHandle)
		user.DELETE("/:userId", handle.DeleteUser)
		user.GET("/:userId/activity", handle.GetUserActivity)
		user.POST("/:userId/follow", handle.FollowUser)
		user.GET("/:userId/followers", handle.ListFollowers)
		user.GET(":userId/statics", handle.GetUserStatics)
	}

	return router
}
