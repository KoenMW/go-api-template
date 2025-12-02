package ports

import service "go-api/domain/service"

type HelloWorldProducer interface {
	Log(msg service.Message) error
}
