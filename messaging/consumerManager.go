package messaging

// ConsumerManager is an interface to consume queue messages
type ConsumerManager interface {
	Consume(queue string)
}
