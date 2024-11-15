package config

import "github.com/rabbitmq/amqp091-go"

func connectRabbitmq() error {
	var err error
	rabbitmq, err = amqp091.Dial(rabbitmqUrl)
	if err != nil {
		rabbitmq.Close()
	}
	return err
}
