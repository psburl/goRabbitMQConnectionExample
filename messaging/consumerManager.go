package messaging

type ConsumerManager interface {
	Consume(queue string)
}
