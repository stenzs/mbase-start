package messageBroker

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func SendMessage(message string) error {

	topic := "my-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:29092", topic, partition)
	if err != nil {
		return err
	}

	err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return err
	}

	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(message)},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		return err
	}

	err = conn.Close()

	if err != nil {
		return err
	}
	return nil
}
