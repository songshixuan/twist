package kfk

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

//KfkSend: send test msg to server
func KfkSend(server string, topic string, partition int) error {
	conn, e := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if e != nil {
		fmt.Println(e.Error())
		return e
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, e = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if e != nil {
		fmt.Println(e.Error())
		return e
	}
	fmt.Println("try read")
	batch := conn.ReadBatch(0, 1e6) // fetch 0 min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}

	batch.Close()

	conn.Close()
	return nil
}
