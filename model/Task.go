package model

import (
	"time"

	"github.com/google/uuid"
)

// https://v1.gorm.io/docs/models.html#Declaring-Models
type Task struct {
	ID          uint `gorm:"primarykey;AUTO_INCREMENT"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time `gorm:"index"`
	Uuid        string    `gorm:"type:varchar(36);unique_index"`
	Description string    `gorm:"type:varchar(512)"`
	DueAt       time.Time
}

func CreateTaskByValues(Description string, DueAt time.Time) *Task {
	newUuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return &Task{
		Description: Description,
		DueAt:       DueAt,
		Uuid:        newUuid.String(),
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
