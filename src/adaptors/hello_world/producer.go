package helloworld

import service "go-api/domain/service"

type HelloWorldProducer struct{}

func (h *HelloWorldProducer) Log(msg service.Message) error {
	println("Logged message:", msg.Message)
	return nil
}
