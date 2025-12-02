package domain

import "fmt"

type Message struct {
	Message string `json:"message"`
}

func (m Message) Validate() error {
	if m.Message == "" {
		return fmt.Errorf("missing message")
	}

	return nil
}
