package helloworld

import "go-api/domain"

type HelloWorldProducer struct{}

func (h *HelloWorldProducer) Log(msg domain.Message) error {
	println("Logged message:", msg.Message)
	return nil
}
