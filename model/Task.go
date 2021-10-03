package model

import (
	"errors"
	"fmt"
	"time"
)

var (
	newTaskError = errors.New("Could not create new task.")
)

// https://v1.gorm.io/docs/models.html#Declaring-Models
type Task struct {
	Identity
	CreatedMetaInfo
	UpdatedMetaInfo
	DeletedMetaInfo

	Description string     `gorm:"type:varchar(512)" json:"description"`
	DueAt       time.Time  `json:"dueAt"`
}

func CreateTaskByValues(newUuid []byte, description string, dueAt time.Time) (*Task, error) {
	newIdentity , err := GetNewIdentity(newUuid)
	if err != nil {
		fmt.Println(newTaskError.Error())
		fmt.Println(err)
		return nil, err
	}

	return &Task{
		Identity: *newIdentity,
		Description: description,
		DueAt:       dueAt,
	}, nil
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
