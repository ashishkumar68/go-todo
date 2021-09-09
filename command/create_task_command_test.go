package command

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateTaskCommandByStreamContent(t *testing.T) {
	content := []byte(`{"uuid":"8419c611-d384-4e84-8114-23a05328e23b","description":"This is a new Task","dueAt":"2021-10-02T15:04:05+07:00"}`)

	command, err := CreateTaskCommandByStreamContent(content)

	assert.Nil(t, err)
	assert.IsType(t, &CreateTaskCommand{}, command)
	assert.Equal(t, "8419c611-d384-4e84-8114-23a05328e23b", command.Uuid)
	assert.Equal(t, "This is a new Task", command.Description)
	assert.Equal(t, "2021-10-02T15:04:05+07:00", command.DueAt.Format(time.RFC3339))
}

func TestCreateTaskCommandByStreamContentWithoutUuid(t *testing.T) {
	content := []byte(`{"description":"This is a new Task","dueAt":"2021-10-02T15:04:05+07:00"}`)

	command, err := CreateTaskCommandByStreamContent(content)

	assert.Nil(t, err)
	assert.IsType(t, &CreateTaskCommand{}, command)
	assert.NotEmpty(t, command.Uuid)
	assert.Equal(t, "This is a new Task", command.Description)
	assert.Equal(t, "2021-10-02T15:04:05+07:00", command.DueAt.Format(time.RFC3339))
}
