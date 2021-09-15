package handler

import (
	"context"
	"sync"

	"github.com/ashishkumar68/go-todo/command"
	"github.com/ashishkumar68/go-todo/database"
	"github.com/ashishkumar68/go-todo/model"
	"github.com/gogolfing/cbus"
	"gorm.io/gorm"
)

type CreateTaskHandler struct {
	manager *gorm.DB
}

func GetTaskHandler() *CreateTaskHandler {
	var once sync.Once
	var createTaskHandler *CreateTaskHandler
	once.Do(func() {
		db, err := database.GetManager()
		if err != nil {
			panic("could not create new DB manager for create new task handler")
		}
		createTaskHandler = &CreateTaskHandler{manager: db}
	})

	return createTaskHandler
}

func (handler CreateTaskHandler) Handle(ctx context.Context, c cbus.Command) (interface{}, error) {
	cmd := c.(*command.CreateTaskCommand)

	task := model.CreateTaskByValues(cmd.Uuid, cmd.Description, cmd.DueAt)
	handler.manager.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(task).Error; err != nil {
			return err
		}
		return nil
	})

	return task, nil
}
