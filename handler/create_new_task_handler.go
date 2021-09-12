package handler

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ashishkumar68/go-todo/command"
	"github.com/ashishkumar68/go-todo/database"
	"github.com/ashishkumar68/go-todo/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	manager, err := database.GetManager()
	if err != nil {
		panic("could not create new DB manager for create new task handler")
	}
	db = manager
}

type CreateTaskHandler struct {
	EventBus *cqrs.EventBus
}

func (handler CreateTaskHandler) HandlerName() string {
	return "CreateTaskHandler"
}

func (handler CreateTaskHandler) NewCommand() interface{} {
	return &command.CreateTaskCommand{}
}

func (handler CreateTaskHandler) Handle(ctx context.Context, c interface{}) error {
	cmd := c.(*command.CreateTaskCommand)

	task := model.CreateTaskByValues(cmd.Uuid, cmd.Description, cmd.DueAt)
	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(task).Error; err != nil {
			return err
		}
		return nil
	})

	return nil
}
