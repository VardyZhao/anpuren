package main

import (
	"gokafka/config"
	"gokafka/consumer"
	"log"
)

func main() {

	log.Println("Application starting...")

	// 加载配置文件
	log.Println("Config file loading...")
	c, err := config.Load("config.yaml")
	if err != nil {
		log.Println("Error loading config.", err)
	}
	log.Println("Config file loaded")

	// todo 载入日志格式

	// 逐个启动消费者，消费者的消费逻辑
	log.Println("Consumer starting...")
	consumer.LaunchConsumer(&c.Kafka)
}
