package model

import (
	"time"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Task struct {
	gorm.Model
	Uuid        string
	Description string
	DueAt       time.Time
}

func CreateTaskByValues(Description string, DueAt time.Time) *Task {
	return &Task{
		Description: Description,
		DueAt:       DueAt,
		Uuid:        uuid.NewString(),
	}
}

func (T *Task) GetDescription() string {
	return T.Description
}

func (T *Task) SetDescription(Description string) *Task {
	T.Description = Description
	return T
}

func (T *Task) GetDueAt() time.Time {
	return T.DueAt
}

func (T *Task) SetDueAt(DueAt time.Time) *Task {
	T.DueAt = DueAt
	return T
}
