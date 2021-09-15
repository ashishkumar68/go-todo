package model

import (
	"time"

	"github.com/google/uuid"
)

// https://v1.gorm.io/docs/models.html#Declaring-Models
type Task struct {
	ID          uint       `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `gorm:"index;default:null" json:"deletedAt,omitempty"`
	Uuid        string     `gorm:"type:varchar(36);unique_index" json:"uuid"`
	Description string     `gorm:"type:varchar(512)" json:"description"`
	DueAt       time.Time  `json:"dueAt"`
}

func CreateTaskByValues(newUuid string, description string, dueAt time.Time) *Task {
	if newUuid == "" {
		nUuid, err := uuid.NewRandom()
		if err != nil {
			panic(err)
		}
		newUuid = nUuid.String()
	}

	return &Task{
		Description: description,
		DueAt:       dueAt,
		Uuid:        newUuid,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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
