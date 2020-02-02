package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

//install lib(match versition) config
func main() {
	topic := "my-topic2"
	partition := 0

	conn, e := kafka.DialLeader(context.Background(), "tcp", "localhost:32770", topic, partition)
	if e != nil {
		fmt.Println(e.Error())
	}
	//conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	fmt.Println("writes")
	i, e := conn.Write(
		[]byte("one!"),
	)
	fmt.Println(i, e)
	conn.Close()
}
