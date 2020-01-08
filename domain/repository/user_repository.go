package repository

import (
    "rest_api/domain/model"
)

type UserRepository interface {
    Insert(DB firestore, userID, name, email string) error
    GetByUserID(DB firestore, userID string) ([]*domain.User, error)
}
