package handler

import (
	"github.com/IBM/sarama"
	"log"
)

type Test1Handler struct {
	BaseHandler
}

func (h *Test1Handler) ProcessMessage(msg *sarama.ConsumerMessage) error {
	log.Println("Handling message in Topic1Handler:", msg)
	return nil
}
