package employee

import "github.com/google/uuid"

type Repository interface {
	GetAll() ([]*Employee, error)
	Add(e *Employee) error
	Delete(UUID uuid.UUID) error
}
