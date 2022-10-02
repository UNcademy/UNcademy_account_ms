package configs

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func RabbitmqConnection() *amqp.Channel {
	urn := fmt.Sprintf("amqp://guest:guest@%s:%s/", os.Getenv("RABBITMQ_CONFIG_HOST"), os.Getenv("RABBITMQ_CONFIG_PORT"))
	conn, err := amqp.Dial(urn)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		logrus.Error(err)
	}

	q, err := ch.QueueDeclare(
		"shistory_history_q",
		false,
		false,
		false,
		false,
		nil,
	)

	logrus.Debug(q)

	if err != nil {
		logrus.Error(err)
	}

	return ch
}
