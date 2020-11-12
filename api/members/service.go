package members

import (
	"amberovsky/teapicker/pkg/member"
	"github.com/google/uuid"
)

/*
Service level for teams composition management
*/

type Service struct {
	repository member.Repository
}

func NewService(repository member.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetTeamMembers(teamUUID uuid.UUID) []*member.Member {
	return s.repository.GetAllForTeam(teamUUID)
}

func (s *Service) Get(UUID uuid.UUID) *member.Member {
	return s.repository.Get(UUID)
}

func (s *Service) Add(teamUUID string, employeeUUID string) error {
	return s.repository.Add(member.NewMember(uuid.New(), uuid.MustParse(teamUUID), uuid.MustParse(employeeUUID)))
}

func (s *Service) Delete(uuidString string) error {
	return s.repository.Delete(uuid.MustParse(uuidString))
}
