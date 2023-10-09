package routes

import (
	"github.com/aditya3232/gatewatchApp-services.git/config"
	"github.com/aditya3232/gatewatchApp-services.git/connection"
	"github.com/aditya3232/gatewatchApp-services.git/handler"
	"github.com/aditya3232/gatewatchApp-services.git/middleware"
	"github.com/aditya3232/gatewatchApp-services.git/model/publisher_vandal_detection"
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	// Initialize repositories
	publisherVandalDetectionRepository := publisher_vandal_detection.NewRepository(connection.RabbitMQ())

	// Initialize services
	publisherVandalDetectionService := publisher_vandal_detection.NewService(publisherVandalDetectionRepository)

	// Initialize handlers
	publisherVandalDetectionHandler := handler.NewPublisherVandalDetectionHandler(publisherVandalDetectionService)

	// Configure routes
	api := router.Group("/publisher/atmvideopack/v1")

	vandalDetectionRoutes := api.Group("/vandaldetection", middleware.ApiKeyMiddleware(config.CONFIG.API_KEY))

	configureVandalDetectionRoutes(vandalDetectionRoutes, publisherVandalDetectionHandler)

}

func configureVandalDetectionRoutes(api *gin.RouterGroup, handler *handler.PublisherVandalDetectionHandler) {
	api.POST("/create", handler.CreateQueueVandalDetection)
}
