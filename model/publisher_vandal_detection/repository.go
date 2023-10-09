package publisher_vandal_detection

import (
	"context"
	"encoding/json"

	"github.com/aditya3232/gatewatchApp-services.git/connection"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Repository interface {
	CreateQueueVandalDetection(input VandalDetectionInput) (VandalDetection, error)
}

type repository struct {
	rabbitmq *amqp.Connection
}

func NewRepository(rabbitmq *amqp.Connection) *repository {
	return &repository{rabbitmq}
}

func (r *repository) CreateQueueVandalDetection(input VandalDetectionInput) (VandalDetection, error) {
	ch, err := connection.RabbitMQ().Channel()
	if err != nil {
		return VandalDetection{}, err
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
		return VandalDetection{}, err
	}

	// body from input
	var inputReadytoMarshal = VandalDetectionInput{
		Tid:           input.Tid,
		DateTime:      input.DateTime,
		Person:        input.Person,
		ConvertedFile: input.ConvertedFile,
	}

	// Convert the VandalDetection struct to JSON
	body, err := json.Marshal(inputReadytoMarshal)
	if err != nil {
		return VandalDetection{}, err
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
		return VandalDetection{}, err
	}

	// Return the sent VandalDetection struct
	return VandalDetection{}, err
}
