package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	description string
}

func (t Task) getDescription() string {
	return t.description
}

func (t *Task) setDescription(description string) *Task {
	t.description = description
	return t
}
