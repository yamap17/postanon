package repository

import (
    "rest_api/domain/model"
)

type UserRepository interface {
    Insert(DB, userID, name, email string) error
    GetByUserID(DB, userID string) ([]*domain.User, error)
}
