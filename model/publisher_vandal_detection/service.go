package publisher_vandal_detection

import "github.com/aditya3232/gatewatchApp-services.git/helper"

type Service interface {
	CreateQueueVandalDetection(input VandalDetectionInput) (VandalDetection, error)
}

type service struct {
	vandalDetectionRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// public message to rmq
func (s *service) CreateQueueVandalDetection(input VandalDetectionInput) (VandalDetection, error) {
	// convert img to base64 with helper.ConvertFileToBase64()
	imgBase64String, err := helper.ConvertFileToBase64(input.File)
	if err != nil {
		return VandalDetection{}, err
	}

	// create vandal detection
	input.ConvertedFile = imgBase64String

	_, err = s.vandalDetectionRepository.CreateQueueVandalDetection(input)
	if err != nil {
		return VandalDetection{}, err
	}

	return VandalDetection{
		Tid:           input.Tid,
		DateTime:      input.DateTime,
		Person:        input.Person,
		ConvertedFile: input.File.Filename,
	}, nil
}
