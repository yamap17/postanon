package repository

import (
	"context"

	"kanban/domain/model"
)

type TaskRepository interface {
	GetAll(context.Context) ([]*model.Task, error)
}