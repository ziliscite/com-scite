package service

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ziliscite/micro-auth/token/internal/domain"
)

type MailPublisher interface {
	SendMail(mail domain.Mail) error
}

type publisher struct {
	amc   *amqp.Connection
	queue amqp.Queue
}

func NewPublisher(amc *amqp.Connection) (MailPublisher, error) {
	// Create a new channel for the publisher
	ch, err := amc.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()

	// Declare the queue once during publisher initialization
	q, err := ch.QueueDeclare(
		"send_mail",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &publisher{
		amc:   amc,
		queue: q,
	}, nil
}

func (p *publisher) SendMail(mail domain.Mail) error {
	ch, err := p.amc.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	body, err := json.Marshal(mail)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",
		p.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
