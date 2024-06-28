package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type ConsumerConfig struct {
	Topic              string        `yaml:"topic"`
	Name               string        `yaml:"name"`
	Group              string        `yaml:"group"`
	RetryTopic         string        `yaml:"retryTopic"`
	RetryMaxTimes      int           `yaml:"retryMaxTimes"`
	ProcessTimeout     int           `yaml:"processTimeout"`
	DeadLetterTopic    string        `yaml:"deadLetterTopic"`
	Handler            string        `yaml:"handler"`
	Partitions         int           `yaml:"partitions"`
	AutoCommit         bool          `yaml:"autoCommit"` // 在批处理后提交
	AutoCommitInterval time.Duration `yaml:"autoCommitInterval"`
}

type KafkaConfig struct {
	BootstrapServers string           `yaml:"bootstrapServers"`
	Consumers        []ConsumerConfig `yaml:"consumers"`
}

type Config struct {
	AppName string      `yaml:"appName"`
	Env     string      `yaml:"env"`
	Version string      `yaml:"version"`
	Kafka   KafkaConfig `yaml:"kafka"`
}

func Load(fileName string) (*Config, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
