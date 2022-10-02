package utils

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func SendMessage(ch *amqp.Channel, username string, program string) {
	err := ch.Publish(
		"",
		"shistory_history_q",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(fmt.Sprintf("{\"Username\": \"%s\", \"Program\": \"%s\"}", username, program)),
		},
	)

	if err != nil {
		logrus.Error(err)
	}
}
