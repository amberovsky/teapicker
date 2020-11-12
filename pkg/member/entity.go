package member

import "github.com/google/uuid"

type Member struct {
	UUID         uuid.UUID `json:"uuid"`
	TeamUUID     uuid.UUID `json:"team_uuid"`
	EmployeeUUID uuid.UUID `json:"employee_uuid"`
}

func NewMember(UUID uuid.UUID, TeamUUID uuid.UUID, EmployeeUUID uuid.UUID) *Member {
	return &Member{
		UUID:         UUID,
		TeamUUID:     TeamUUID,
		EmployeeUUID: EmployeeUUID,
	}
}
