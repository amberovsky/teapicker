package pick

import (
	"amberovsky/teapicker/pkg/pick"
	"github.com/google/uuid"
)

/*
Service level for picking next tea maker
*/

type Service struct {
	strategy pick.Strategy
}

func NewService(strategy pick.Strategy) *Service {
	return &Service{
		strategy: strategy,
	}
}

func (s *Service) Pick(teamUuidString string) (*uuid.UUID, error) {
	return s.strategy(uuid.MustParse(teamUuidString))
}
