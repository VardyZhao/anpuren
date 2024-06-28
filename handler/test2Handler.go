package handler

import (
	"github.com/IBM/sarama"
	"log"
)

type Test2Handler struct {
	BaseHandler
}

func (h *Test2Handler) ProcessMessage(msg *sarama.ConsumerMessage) error {
	log.Println("Handling message in Topic2Handler:", msg)
	return nil
}
