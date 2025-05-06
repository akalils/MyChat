package main

import (
	"LishengChat-go/internal/config"
	"LishengChat-go/internal/https_server"
	"LishengChat-go/internal/service/chat"
	"LishengChat-go/internal/service/kafka"
	myredis "LishengChat-go/internal/service/redis"
	"LishengChat-go/pkg/zlog"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf := config.GetConfig()
	host := conf.MainConfig.Host
	port := conf.MainConfig.Port
	kafkaConfig := conf.KafkaConfig
	if kafkaConfig.MessageMode == "kafka" {
		kafka.KafkaService.KafkaInit()
	}
	if kafkaConfig.MessageMode == "channel" {
		go chat.ChatServer.Start()
	} else {
		go chat.KafkaChatServer.Start()
	}

	go func () {
		if err := https_server.GE.RunTLS(fmt.Sprintf("%s:%d", host, port), "/etc/ssl/certs/server.crt", "/etc/ssl/private/server.key"); err != nil {
			zlog.Fatal("server running fault")
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	if kafkaConfig.MessageMode == "kafka" {
		kafka.KafkaService.KafkaClose()
	}
	chat.ChatServer.Close()

	zlog.Info("关闭服务器...")
	if err := myredis.DeleteAllRedisKeys(); err != nil {
		zlog.Error(err.Error())
	} else {
		zlog.Info("所有Redis键已删除")
	}
	zlog.Info("服务器已关闭")
}