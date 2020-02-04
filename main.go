package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/songshixuan/twist/pkg/kfk"
)

//install lib(match versition) config
func main() {
	topic := "my-topic2"
	partition := 0
	kfk.KfkSend("127.0.0.1:9092", topic, partition)
	fmt.Println("read--")
	//read
	conn, e := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if e != nil {
		panic(e)
	}
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

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
}
