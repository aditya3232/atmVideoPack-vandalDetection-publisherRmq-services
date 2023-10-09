package handler

import (
	"net/http"

	"github.com/aditya3232/gatewatchApp-services.git/constant"
	"github.com/aditya3232/gatewatchApp-services.git/helper"
	"github.com/aditya3232/gatewatchApp-services.git/log"
	"github.com/aditya3232/gatewatchApp-services.git/model/publisher_vandal_detection"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type PublisherVandalDetectionHandler struct {
	publisherVandalDetectionService publisher_vandal_detection.Service
}

func NewPublisherVandalDetectionHandler(publisherVandalDetectionService publisher_vandal_detection.Service) *PublisherVandalDetectionHandler {
	return &PublisherVandalDetectionHandler{publisherVandalDetectionService}
}

// public message to rmqg
func (h *PublisherVandalDetectionHandler) CreateQueueVandalDetection(c *gin.Context) {
	var input publisher_vandal_detection.VandalDetectionInput

	err := c.ShouldBindWith(&input, binding.FormMultipart)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.InvalidRequest, http.StatusBadRequest, errorMessage)
		log.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	vandalDetection, err := h.publisherVandalDetectionService.CreateQueueVandalDetection(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, errorMessage)
		log.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessMessage, http.StatusOK, publisher_vandal_detection.FormatPublisherVandalDetection(vandalDetection))
	log.Info("Queue vandal detection berhasil dibuat")
	c.JSON(http.StatusOK, response)
}
