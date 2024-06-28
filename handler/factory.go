package handler

import (
	"gokafka/config"
	"log"
)

// 定义一个map来映射字符串和Handler
var handlerMap = map[string]func(c config.ConsumerConfig) Handler{
	"test1Handler": func(c config.ConsumerConfig) Handler { return &Test1Handler{BaseHandler{Config: c}} },
	"test2Handler": func(c config.ConsumerConfig) Handler { return &Test2Handler{BaseHandler{Config: c}} },
}

// GetHandler 工厂方法，根据字符串返回对应的Handler实例
func GetHandler(c config.ConsumerConfig) Handler {
	if handlerFunc, exists := handlerMap[c.Handler]; exists {
		return handlerFunc(c)
	}

	log.Fatalf("handler not found: %s", c.Handler)
	return nil
}
