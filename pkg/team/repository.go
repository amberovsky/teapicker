package team

import "github.com/google/uuid"

type Repository interface {
	GetAll() ([]*Team, error)
	Add(t *Team) error
	Delete(UUID uuid.UUID) error
}
