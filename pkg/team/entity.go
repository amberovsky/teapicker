package team

import (
	"github.com/google/uuid"
)

type Team struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

func NewTeam(UUID uuid.UUID, Name string) *Team {
	return &Team{
		UUID: UUID,
		Name: Name,
	}
}
