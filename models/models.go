package models

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Uuid        uuid.UUID
	PublishedAt time.Time
	Airac       int64
	Files       []string
}
