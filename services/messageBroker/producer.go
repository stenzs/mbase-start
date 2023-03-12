package messageBroker

import (
	"encoding/json"

	"github.com/adjust/rmq"

	"mbase/models"
)

type Task struct {
	id    int
	name  string
	name2 string
}

func SendMessage(message models.Task) error {
	connection := rmq.OpenConnection("mbase", "tcp", "localhost:6379", 1)
	taskQueue := connection.OpenQueue("tasks")

	taskBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	taskQueue.PublishBytes(taskBytes)

	return nil

}
