package models

import (
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}
