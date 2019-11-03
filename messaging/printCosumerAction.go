package messaging

import "fmt"

// PrintConsumerAction implements ConsumerAction printing received messages
type PrintConsumerAction struct {
}

// NewPrintConsumerAction constructs a new *PrintConsumerAction
func NewPrintConsumerAction() *PrintConsumerAction {
	return &PrintConsumerAction{}
}

// Execute prints a received message
func (*PrintConsumerAction) Execute(message []byte) {
	fmt.Println(string(message))
}
