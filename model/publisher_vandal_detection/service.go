package publisher_vandal_detection

import (
	"github.com/aditya3232/atmVideoPack-vandalDetection-publisherRmq-services.git/helper"
	"github.com/aditya3232/atmVideoPack-vandalDetection-publisherRmq-services.git/model/tb_tid"
)

type Service interface {
	/*
		- input ada 2, pertama input dari API
		- kedua input yang akan dimasukkan ke RMQ
		- disini returnnya yg akan ditampilkan di API adalah inputan rmq,
		- disini parameter adalah input
	*/
	CreateQueueVandalDetection(input RmqPublisherVandalDetectionInput) (RmqPublisherVandalDetection, error)
}

type service struct {
	vandalDetectionRepository Repository
	tbTidRepository           tb_tid.Repository
}

func NewService(repository Repository, tbTidRepository tb_tid.Repository) *service {
	return &service{repository, tbTidRepository}
}

// public message to rmq
func (s *service) CreateQueueVandalDetection(input RmqPublisherVandalDetectionInput) (RmqPublisherVandalDetection, error) {
	var rmqPublisherVandalDetection RmqPublisherVandalDetection

	// validasi is image
	err := helper.IsImage(input.FileCaptureVandalDetection)
	if err != nil {
		return rmqPublisherVandalDetection, err
	}

	// convert image to jpg
	err = helper.ConvertImageToJpg(input.FileCaptureVandalDetection)
	if err != nil {
		return rmqPublisherVandalDetection, err
	}

	// convert img
	imgBase64String, err := helper.ConvertFileToBase64(input.FileCaptureVandalDetection)
	if err != nil {
		return rmqPublisherVandalDetection, err
	}

	// get name file
	fileName := input.FileCaptureVandalDetection.Filename

	// get id from input tid
	tidID, err := s.tbTidRepository.GetOneByTid(input.Tid)
	if err != nil {
		return rmqPublisherVandalDetection, err
	}

	newRmqPublisherVandalDetection := RmqPublisherVandalDetection{
		TidID:                               &tidID.ID,
		Tid:                                 input.Tid,
		DateTime:                            input.DateTime,
		Person:                              input.Person,
		ConvertedFileCaptureVandalDetection: imgBase64String,
		FileNameCaptureVandalDetection:      fileName,
	}

	newVandalDetection, err := s.vandalDetectionRepository.CreateQueueVandalDetection(newRmqPublisherVandalDetection)
	if err != nil {
		return newVandalDetection, err
	}

	return newVandalDetection, nil
}
