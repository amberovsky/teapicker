package teams

import (
	"amberovsky/teapicker/pkg/member"
	"amberovsky/teapicker/pkg/team"
	"github.com/google/uuid"
)

/*
Service level for teams management
*/

type Service struct {
	repository       team.Repository
	memberRepository member.Repository
}

func NewService(repository team.Repository, memberRepository member.Repository) *Service {
	return &Service{
		repository:       repository,
		memberRepository: memberRepository,
	}
}

func (s *Service) GetAll() ([]*team.Team, error) {
	return s.repository.GetAll()
}

func (s *Service) Add(name string) error {
	return s.repository.Add(team.NewTeam(uuid.New(), name))
}

func (s *Service) Delete(uuidString string) error {
	teamUUID := uuid.MustParse(uuidString)

	_ = s.repository.Delete(teamUUID)
	_ = s.memberRepository.DeleteAllFromTeam(teamUUID)

	return nil
}
