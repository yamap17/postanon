package usecase

import (
	"context"

	"kanban/domain/model"
	"kanban/domain/repository"
)

type TaskUseCase interface {
	GetAll(context.Context) ([]*model.Task, error)
}

type taskUseCase struct {
	taskRepository repository.TaskRepository
}

func NewTaskUseCase(tr repository.TaskRepository) TaskUseCase {
	return &taskUseCase{
		taskRepository: tr,
	}
}

func (tu taskUseCase) GetAll(ctx context.Context) (tasks []*model.Task, err error) {
	tasks, err = tu.taskRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
