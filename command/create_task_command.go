package command

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type CreateTaskCommand struct {
	Uuid        string    `json:"uuid"`
	Description string    `validate:"required" json:"description"`
	DueAt       time.Time `validate:"required" json:"dueAt"`
}

func CreateTaskCommandByStreamContent(data []byte) (*CreateTaskCommand, error) {
	var command CreateTaskCommand
	err := json.Unmarshal(data, &command)
	if err != nil {
		return nil, errors.New("Could not decode the stream while creating CreateTaskCommand")
	}
	if command.Uuid == "" {
		taskUuid, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		command.Uuid = taskUuid.String()
	}

	return &command, nil
}

func CreateTaskCommandByValues(
	description string,
	dueAt time.Time,
	taskUuid string,
) (*CreateTaskCommand, error) {
	var newUuid string
	if taskUuid != "" {
		newUuid = taskUuid
	} else {
		taskUuid, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		newUuid = taskUuid.String()
	}

	return &CreateTaskCommand{
		Uuid:        newUuid,
		Description: description,
		DueAt:       dueAt,
	}, nil
}
