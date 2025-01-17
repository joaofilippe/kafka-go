package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "my-go-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9093", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	defer conn.Close()

	kafkaChann := make(chan string)

	go func() {
		fmt.Println("Enter message to send to Kafka:")
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			kafkaChann <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Println("error reading from console:", err)
		}

		close(kafkaChann)
	}()

	for message := range kafkaChann {
		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		_, err := conn.WriteMessages(
			kafka.Message{Value: []byte(message)},
		)

		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

		fmt.Println("\nsent message:", message)
		fmt.Println("\nEnter message to sendo to Kafka:")
	}
}
