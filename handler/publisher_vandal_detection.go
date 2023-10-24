package handler

import (
	"net/http"

	"github.com/aditya3232/atmVideoPack-vandalDetection-publisherRmq-services.git/constant"
	"github.com/aditya3232/atmVideoPack-vandalDetection-publisherRmq-services.git/helper"
	log_function "github.com/aditya3232/atmVideoPack-vandalDetection-publisherRmq-services.git/log"
	"github.com/aditya3232/atmVideoPack-vandalDetection-publisherRmq-services.git/model/publisher_vandal_detection"
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
	var input publisher_vandal_detection.RmqPublisherVandalDetectionInput

	err := c.ShouldBindWith(&input, binding.FormMultipart)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.InvalidRequest, http.StatusBadRequest, errorMessage)
		log_function.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	vandalDetection, err := h.publisherVandalDetectionService.CreateQueueVandalDetection(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, errorMessage)
		log_function.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessMessage, http.StatusOK, publisher_vandal_detection.PublisherVandalDetectionFormat(vandalDetection))
	log_function.Info("Queue vandal detection berhasil dibuat")
	c.JSON(http.StatusOK, response)
}
