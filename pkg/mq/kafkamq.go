package mq

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

type (
	Kafka struct {
		producer    *kafka.Producer
		consumer    *kafka.Consumer
		brokers     string // Kafka连接地址
		adminClient *kafka.AdminClient
		Done        chan bool
		IsReady     bool
	}

	// KafkaConfig Kafka配置
	KafkaConfig struct {
		Brokers string
		GroupID string
	}
)

// NewKafka 创建一个新的Kafka实例，并自动尝试连接到服务器
func NewKafka(config *KafkaConfig) *Kafka {
	kafkaClient := &Kafka{brokers: config.Brokers, Done: make(chan bool, 1)}
	err := kafkaClient.connect()
	if err != nil {
		fmt.Printf("Failed to connect to Kafka: %v\n", err)
		return nil
	}
	go kafkaClient.checkAndReconnect()
	return kafkaClient
}

// connect 创建一个新的Kafka管理客户端
func (k *Kafka) connect() error {
	adminClient, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": k.brokers})
	if err != nil {
		return fmt.Errorf("failed to create admin client: %w", err)
	}
	k.adminClient = adminClient
	// 创建生产者配置
	producerConfig := &kafka.ConfigMap{
		"bootstrap.servers": k.brokers,
	}
	// 创建生产者
	producer, err := kafka.NewProducer(producerConfig)
	if err != nil {
		return fmt.Errorf("创建生产者失败: %w", err)
	}
	// 创建生产者配置
	consumerConfig := &kafka.ConfigMap{
		"bootstrap.servers": k.brokers,
	}
	// 使用外部传入的消费者配置
	consumerConfig.SetKey("bootstrap.servers", k.brokers)  // 确保包含 brokers 配置
	consumerConfig.SetKey("group.id", "test-group")        // 确保包含 brokers 配置
	consumerConfig.SetKey("auto.offset.reset", "earliest") // 确保包含 brokers 配置

	// 创建消费者
	consumer, err := kafka.NewConsumer(consumerConfig)
	if err != nil {
		producer.Close() // 在创建消费者失败时关闭生产者
		return fmt.Errorf("创建消费者失败: %w", err)
	}
	k.producer = producer
	k.consumer = consumer
	return nil
}

// checkAndReconnect 检查连接状态并自动重连
func (k *Kafka) checkAndReconnect() {
	maxRetries := 5
	retryInterval := 5 * time.Second
	for {
		_, err := k.adminClient.GetMetadata(nil, false, 60000)
		if err != nil {
			fmt.Printf("Connection to Kafka lost: %v\n", err)
			k.adminClient.Close()
			k.adminClient = nil
			for attempt := 0; attempt < maxRetries; attempt++ {
				err := k.connect()
				if err != nil {
					fmt.Printf("Failed to reconnect to Kafka on attempt %d: %v\n", attempt+1, err)
					time.Sleep(retryInterval)
					continue
				}
				fmt.Println("Successfully reconnected to Kafka")
				break
			}
			if k.adminClient == nil {
				fmt.Printf("Failed to reconnect to Kafka after %d attempts\n", maxRetries)
			}
		} else {
			k.IsReady = true
		}
		time.Sleep(10 * time.Second) // 每10秒检查一次连接状态
	}
}

// Push 将数据推送到Kafka主题中
func (k *Kafka) Push(topic string, data []byte) error {
	if !k.IsReady {
		return errNotConnected
	}
	k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)
	return nil
}

// Consume 从Kafka主题中消费数据
func (k *Kafka) Consume(topic string) (<-chan *kafka.Message, error) {
	if !k.IsReady {
		return nil, errNotConnected
	}
	if err := k.consumer.Subscribe(topic, nil); err != nil {
		return nil, err
	}

	msgChan := make(chan *kafka.Message)
	go func() {
		for {
			select {
			case _ = <-k.Done:
				close(msgChan)
				return
			default:
				msg, err := k.consumer.ReadMessage(-1)
				if err == nil {
					msgChan <- msg
				}
			}
		}
	}()
	return msgChan, nil
}

// Close 关闭生产者和消费者。
func (k *Kafka) Close() error {
	if !k.IsReady {
		return errAlreadyClosed
	}
	k.Done <- true
	close(k.Done)
	k.producer.Close()
	k.consumer.Close()
	k.IsReady = false
	return nil
}
