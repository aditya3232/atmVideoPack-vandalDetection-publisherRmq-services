package publisher_vandal_detection

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Repository interface {
	/*
		- ambil input dari API, dimasukkan ke entity rmq
		- disini parameter adalah entitiy
		- disini returnnya yg akan ditampilkan di API adalah entity rmq,
	*/
	CreateQueueVandalDetection(rmqPublisherVandalDetection RmqPublisherVandalDetection) (RmqPublisherVandalDetection, error)
}

type repository struct {
	rabbitmq *amqp.Connection
}

func NewRepository(rabbitmq *amqp.Connection) *repository {
	return &repository{rabbitmq}
}

func (r *repository) CreateQueueVandalDetection(rmqPublisherVandalDetection RmqPublisherVandalDetection) (RmqPublisherVandalDetection, error) {

	ch, err := r.rabbitmq.Channel()
	if err != nil {
		return rmqPublisherVandalDetection, err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"VandalDetectionQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return rmqPublisherVandalDetection, err
	}

	// yang dimarshal adalah entity rmq Vandal detection
	inputReadytoMarshal := RmqPublisherVandalDetection{
		Tid:                                 rmqPublisherVandalDetection.Tid,
		DateTime:                            rmqPublisherVandalDetection.DateTime,
		Person:                              rmqPublisherVandalDetection.Person,
		ConvertedFileCaptureVandalDetection: rmqPublisherVandalDetection.ConvertedFileCaptureVandalDetection,
	}

	// Convert the VandalDetection struct to JSON
	body, err := json.Marshal(inputReadytoMarshal)
	if err != nil {
		return rmqPublisherVandalDetection, err
	}

	ctx := context.Background() // Create a context
	err = ch.PublishWithContext(ctx,
		"",
		"VandalDetectionQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return rmqPublisherVandalDetection, err
	}

	return rmqPublisherVandalDetection, err
}
