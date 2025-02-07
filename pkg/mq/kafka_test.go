package mq

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestKafkaConnectPushConsume(t *testing.T) {
	// 假设你有一个可用的Kafka集群，并且配置如下：
	kafkaConfig := &KafkaConfig{
		Brokers: "localhost:9092", // 替换为你的Kafka地址
		GroupID: "test-group",
	}

	// 1. 创建一个新的Kafka实例并确保连接成功
	kafkaClient := NewKafka(kafkaConfig)
	if kafkaClient == nil {
		t.Fatal("Failed to create Kafka client")
	}
	defer func(kafkaClient *Kafka) {
		err := kafkaClient.Close()
		if err != nil {
			t.Fatal(err.Error())
		}
	}(kafkaClient)

	// 确保kafkaClient已经准备好
	for !kafkaClient.IsReady {
		time.Sleep(100 * time.Millisecond)
	}

	// 2. 推送一条消息到指定的主题
	topic := "test-topic"
	message := []byte("Hello, Kafka!")
	err := kafkaClient.Push(topic, message)
	assert.NoError(t, err, "Failed to push message")
	// 3. 消费该主题的消息并验证接收到的消息是否与推送的消息一致
	msgChan, err := kafkaClient.Consume(topic)
	assert.NoError(t, err, "Failed to consume messages")
	var receivedMessage []byte
	for {
		select {
		case msg := <-msgChan:
			receivedMessage = msg.Value
			fmt.Printf(string(receivedMessage))
			return
		case <-time.After(10 * time.Second):
			t.Fatal("Timed out waiting for message")
			return
		}
	}
}
