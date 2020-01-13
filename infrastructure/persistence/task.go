package persistence

import (
	"context"
	"time"

	"kanban/domain/model"
	"kanban/domain/repository"
)

type taskPersistence struct{}

func NewTaskPersistence() repository.TaskRepository {
	return &taskPersistence{}
}

// 簡略化のためにモックデータを返却
func (tp taskPersistence) GetAll(context.Context) ([]*model.Task, error) {
	task1 := model.Task{}
	task1.Id = 1
	task1.Title = "タスク1"
	task1.CreatedAt = time.Now().Add(-24 * time.Hour)

	task2 := model.Task{}
	task2.Id = 1
	task2.Title = "タスク2"
	task2.CreatedAt = time.Now().Add(-24 * time.Hour)

	return []*model.Task{&task1, &task2}, nil
}