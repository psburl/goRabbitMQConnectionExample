package messaging

import (
	env "../environment"
	"github.com/streadway/amqp"
)

type RabbitConsumerManager struct {
	consumerAction ConsumerAction
	environment    env.Manager
}

func NewRabbitConsumerManager(
	environment env.Manager,
	consumerAction ConsumerAction) *RabbitConsumerManager {
	return &RabbitConsumerManager{
		environment:    environment,
		consumerAction: consumerAction}
}

func (consumer *RabbitConsumerManager) Consume(queue string) {

	amqpConn, err := amqp.Dial(consumer.environment.GetAmqpURL())
	errorHandle("Failed to connect to RabbitMQ", err)
	defer amqpConn.Close()

	amqpChan, err := amqpConn.Channel()
	errorHandle("Failed to open a channel", err)
	defer amqpChan.Close()

	messages, err := amqpChan.Consume(queue, "GoLang", false, false, false, false, nil)

	forever := make(chan (bool))

	go func() {
		for message := range messages {
			consumer.consumerAction.Execute(message.Body)
		}
	}()

	<-forever
}

func errorHandle(message string, err error) {
	panic(message + ":" + err.Error())
}
