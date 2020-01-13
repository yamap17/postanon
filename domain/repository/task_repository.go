package repository

import (
	"context"

	"rest_api/domain/model"
)

type TaskRepository interface {
	GetAll(context.Context) ([]*model.Task, error)
}