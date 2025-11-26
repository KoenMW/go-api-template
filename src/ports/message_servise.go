package ports

import "go-api/domain"

type HelloWorldProducer interface {
	Log(msg domain.Message) error
}
