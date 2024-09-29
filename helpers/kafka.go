package helpers

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fomichalopoulos/companiesMicroService/models"
)

func ProdKafkaMsg(msg string, kafkaCfg models.KafkaConfig, producer *kafka.Producer) error {

	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kafkaCfg.KAFKA_TOPIC, Partition: kafka.PartitionAny}, // Topic and partition
		Value:          []byte(msg),                                                                       // The message payload
	}, nil)

	if err != nil {
		fmt.Printf("Failed to produce message: %s\n", err)
		return err
	} else {
		fmt.Println("Message successfully sent to Kafka!")
	}

	// 6. Wait for message deliveries (optional)
	producer.Flush(1000) // Wait for up to 15 seconds for message delivery confirmation
	fmt.Println("passed flush")
	return nil

}
