package model

import "time"

type Task struct {
	Id        int64
	Title     string
	CreatedAt time.Time
}