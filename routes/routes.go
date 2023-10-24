package routes

import (
	"github.com/aditya3232/atmVideoPack-vandalDetection-publisherRmq-services.git/config"
	"github.com/aditya3232/atmVideoPack-vandalDetection-publisherRmq-services.git/connection"
	"github.com/aditya3232/atmVideoPack-vandalDetection-publisherRmq-services.git/handler"
	"github.com/aditya3232/atmVideoPack-vandalDetection-publisherRmq-services.git/middleware"
	"github.com/aditya3232/atmVideoPack-vandalDetection-publisherRmq-services.git/model/publisher_vandal_detection"
	"github.com/aditya3232/atmVideoPack-vandalDetection-publisherRmq-services.git/model/tb_tid"
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	// Initialize repositories
	publisherVandalDetectionRepository := publisher_vandal_detection.NewRepository(connection.RabbitMQ())
	tbTidRepository := tb_tid.NewRepository(connection.DatabaseMysql())

	// Initialize services
	publisherVandalDetectionService := publisher_vandal_detection.NewService(publisherVandalDetectionRepository, tbTidRepository)

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
