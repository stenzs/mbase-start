package messageBroker

import (
	"encoding/json"
	"log"

	"github.com/adjust/rmq/v5"

	"mbase/models"
)

type Task struct {
	id    int
	name  string
	name2 string
}

func SendMessage(message models.Task) error {
	errChan := make(chan error, 100)
	go logErrors(errChan)

	connection, err := rmq.OpenConnection("mbase", "tcp", "localhost:6379", 1, errChan)
	if err != nil {
		return err
	}

	taskQueue, err := connection.OpenQueue("tasks")
	if err != nil {
		return err
	}

	taskBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = taskQueue.PublishBytes(taskBytes)
	if err != nil {
		return err
	}

	return nil
}

func logErrors(errChan <-chan error) {
	for err := range errChan {
		switch err := err.(type) {
		case *rmq.HeartbeatError:
			if err.Count == rmq.HeartbeatErrorLimit {
				log.Print("heartbeat error (limit): ", err)
			} else {
				log.Print("heartbeat error: ", err)
			}
		case *rmq.ConsumeError:
			log.Print("consume error: ", err)
		case *rmq.DeliveryError:
			log.Print("delivery error: ", err.Delivery, err)
		default:
			log.Print("other error: ", err)
		}
	}
}
