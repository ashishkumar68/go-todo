package repository

import (
	"sync"

	"github.com/ashishkumar68/go-todo/database"
	"github.com/ashishkumar68/go-todo/model"
)

var (
	taskRepository *TaskRepository
	once           sync.Once
)

type TaskRepository struct {
	BaseRepository
}

func GetTaskRepository() *TaskRepository {
	once.Do(func() {
		db, err := database.GetManager()
		if err != nil {
			panic(err)
		}
		taskRepository = &TaskRepository{
			BaseRepository{Manager: db},
		}
	})

	return taskRepository
}

func (t *TaskRepository) FindOneById(id int) *model.Task {
	var task *model.Task
	t.Manager.First(&task, id)

	return task
}

func (t *TaskRepository) FindOneByUuid(uuid string) *model.Task {
	var task *model.Task
	t.Manager.First(&task, "uuid = ?", uuid)

	return task
}
