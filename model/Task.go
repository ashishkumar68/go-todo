package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	description string
	dueAt       time.Time
}

func (t Task) GetDescription() string {
	return t.description
}

func (t *Task) SetDescription(description string) *Task {
	t.description = description
	return t
}

func CreateTaskByValues(description string, dueAt time.Time) *Task {
	return &Task{
		description: description,
		dueAt:       dueAt,
	}
}
