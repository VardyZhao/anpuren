package consumer

import (
	"fmt"
	"github.com/IBM/sarama"
	"gokafka/config"
	"gokafka/handler"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

func LaunchConsumer(kc *config.KafkaConfig) {
	var wg sync.WaitGroup
	for _, cc := range kc.Consumers {
		for i := 1; i <= cc.Partitions; i++ {
			wg.Add(1)
			go func(cc config.ConsumerConfig) {
				defer wg.Done()
				createConsumer(kc, cc, i)
			}(cc)
		}
	}

	wg.Wait()
}

func initConsumerConfig(cc config.ConsumerConfig, index int) *sarama.Config {
	consumerName := cc.Name + "-" + strconv.Itoa(index)
	sc := sarama.NewConfig()
	sc.ClientID = consumerName
	sc.Consumer.Offsets.AutoCommit.Enable = cc.AutoCommit
	sc.Consumer.Offsets.AutoCommit.Interval = cc.AutoCommitInterval * time.Second
	return sc
}

func createConsumer(kc *config.KafkaConfig, cc config.ConsumerConfig, index int) {
	// 初始化消费者config
	sc := initConsumerConfig(cc, index)

	// 创建消费者
	consumerGroup, err := sarama.NewConsumerGroup(strings.Split(kc.BootstrapServers, ","), cc.Group, sc)
	if err != nil {
		log.Fatalf("Error creating consumer group: %v", err)
	}

	defer consumerGroup.Close()

	for {
		// 绑定handler，循环消费
		err := consumerGroup.Consume(nil, []string{cc.Topic}, &Consumer{
			handler.GetHandler(cc),
		})
		if err != nil {
			log.Printf("Error in consumer group: %v", err)
		}
	}
}

type Consumer struct {
	Handler handler.Handler
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for message := range claim.Messages() {
		fmt.Printf("Message claimed: value = %s, timestamp = %v, topic = %s\n", string(message.Value), message.Timestamp, message.Topic)
		// 处理消息
		err := consumer.Handler.ProcessMessage(message)
		if err != nil {
			// todo 判断重试次数
			// todo 进入重试队列
			// todo 进入死信队列
		}

		// 标记消息已处理
		session.MarkMessage(message, "")

		// 设置了手动提交，就马上提交
		if consumer.Handler.GetAutoCommit() {
			session.Commit()
		}
	}
	return nil
}
