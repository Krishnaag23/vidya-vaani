package streaming

import (
    "log"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)


type KafkaProducer struct {
    producer *kafka.Producer
}


func NewKafkaProducer(broker string) *KafkaProducer {
    p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
    if err != nil {
        log.Fatalf("Failed to create producer: %s", err)
    }

    return &KafkaProducer{producer: p}
}


func (kp *KafkaProducer) Produce(topic string, message string) {
    
    kafkaMsg := &kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Value:          []byte(message),
    }

    
    err := kp.producer.Produce(kafkaMsg, nil)
    if err != nil {
        log.Printf("Failed to produce message: %s", err)
    }
}


func (kp *KafkaProducer) Close() {
    kp.producer.Flush(15 * 1000) 
    kp.producer.Close()
}


type KafkaConsumer struct {
    consumer *kafka.Consumer
}


func NewKafkaConsumer(broker string, groupID string, topics []string) *KafkaConsumer {
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": broker,
        "group.id":          groupID,
        "auto.offset.reset": "earliest", 
    })
    if err != nil {
        log.Fatalf("Failed to create consumer: %s", err)
    }

    c.SubscribeTopics(topics, nil)

    return &KafkaConsumer{consumer: c}
}


func (kc *KafkaConsumer) Consume() {
    for {
        msg, err := kc.consumer.ReadMessage(-1)
        if err == nil {
            log.Printf("Consumed message: %s", string(msg.Value))
        } else {
            log.Printf("Consumer error: %v", err)
        }
    }
}


func (kc *KafkaConsumer) Close() {
    kc.consumer.Close()
}
