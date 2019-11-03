package messaging

import "fmt"

type PrintConsumerAction struct {
}

func NewPrintConsumerAction() *PrintConsumerAction {
	return &PrintConsumerAction{}
}

func (*PrintConsumerAction) Execute(message []byte) {
	fmt.Println(string(message))
}
