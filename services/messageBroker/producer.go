package messageBroker

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func SendMessage(key, value string) error {

	partition, _ := strconv.Atoi(os.Getenv("KAFKA_CONNECTION_PARTITION"))
	conn, err := kafka.DialLeader(
		context.Background(),
		"tcp",
		os.Getenv("KAFKA_CONNECTION_HOST"),
		os.Getenv("KAFKA_CONNECTION_TOPIC"),
		partition)
	if err != nil {
		return err
	}

	err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return err
	}

	_, err = conn.WriteMessages(
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(value),
		},
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
