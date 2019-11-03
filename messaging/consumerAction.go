package messaging

type ConsumerAction interface {
	Execute(message []byte)
}
