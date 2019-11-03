package messaging

// ConsumerAction is an interface invoked by consumerManager
// to execute an action on received message
type ConsumerAction interface {
	Execute(message []byte)
}
