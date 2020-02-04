package kfk

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Server    string
	Topic     string
	Partition int
}
type KafkaConsumer struct {
	Server    string
	Topic     string
	Partition int
}

func (consumer *KafkaConsumer) Consume() error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{consumer.Server},
		Topic:    consumer.Topic,
		MinBytes: 0,   // 0KB
		MaxBytes: 1e6, // 1MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	r.Close()
	return nil
}

//KfkSend: send test msg to server
func (producer *KafkaProducer) KfkSend(msgs []string) error {
	conn, e := kafka.DialLeader(context.Background(), "tcp",
		producer.Server,
		producer.Topic,
		producer.Partition,
	)
	if e != nil {
		fmt.Println(e.Error())
		return e
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	kafkaMsgs := make([]kafka.Message, 0, len(msgs))
	for _, m := range msgs {
		kafkaMsgs = append(kafkaMsgs, kafka.Message{Value: []byte(m)})
	}
	_, e = conn.WriteMessages(
		kafkaMsgs...,
	)
	if e != nil {
		fmt.Println(e.Error())
		return e
	}

	return nil
}
