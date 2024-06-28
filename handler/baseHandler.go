package handler

import (
	"github.com/IBM/sarama"
	"gokafka/config"
	"log"
)

type Handler interface {
	ProcessMessage(msg *sarama.ConsumerMessage) error
	SendToRetryQueue(msg string) error
	SendToDeadLetterQueue(msg string) error
	GetRetryMaxTimes() int
	GetAutoCommit() bool
}

type BaseHandler struct {
	Config config.ConsumerConfig
}

func (h *BaseHandler) ProcessMessage(msg *sarama.ConsumerMessage) error {
	log.Println("Handling message in BaseHandler:", msg)
	return nil
}

func (h *BaseHandler) SendToRetryQueue(msg string) error {
	log.Println("Send to retry queue:", msg)
	return nil
}

func (h *BaseHandler) SendToDeadLetterQueue(msg string) error {
	log.Println("Send to dead letter queue:", msg)
	return nil
}

func (h *BaseHandler) GetRetryMaxTimes() int {
	return h.Config.RetryMaxTimes
}

func (h *BaseHandler) GetAutoCommit() bool {
	return h.Config.AutoCommit
}
