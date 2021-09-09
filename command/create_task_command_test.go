package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTaskCommandByStreamContent(t *testing.T) {
	content := []byte(`{"uuid":"8419c611-d384-4e84-8114-23a05328e23b","description":"This is a new Task","dueAt":"2021-10-02T15:04:05+07:00"}`)

	command, err := CreateTaskCommandByStreamContent(content)

	assert.Nil(t, err)
	assert.IsType(t, &CreateTaskCommand{}, command)
}
