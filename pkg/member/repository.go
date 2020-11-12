package member

import "github.com/google/uuid"

type Repository interface {
	GetAllForTeam(teamUUID uuid.UUID) []*Member
	Get(UUID uuid.UUID) *Member
	Add(m *Member) error
	Delete(UUID uuid.UUID) error

	DeleteAllFromTeam(teamUUID uuid.UUID) error
	DeleteEmployee(employeeUUID uuid.UUID) error
}
