package main

import (
	"os"

	"github.com/songshixuan/twist/pkg/kfk"
)

//install lib(match versition) config
func main() {

	switch os.Args[1] {
	case "producer":
		producer := kfk.KafkaProducer{
			Server:    "127.0.0.1:9092",
			Topic:     "what-the-fuck",
			Partition: 0,
		}
		producer.KfkSend([]string{"how", "I", "Met", "Your", "Mother"})
	case "c":

		counsumer := kfk.KafkaConsumer{
			Server: "127.0.0.1:9092",
			Topic:  "what-the-fuck",
		}
		counsumer.Consume()
	}

}
