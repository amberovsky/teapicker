package employee

import (
	"github.com/google/uuid"
)

type Employee struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

func NewEmployee(UUID uuid.UUID, Name string) *Employee {
	return &Employee{
		UUID: UUID,
		Name: Name,
	}
}
