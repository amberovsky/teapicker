package employees

import (
	"amberovsky/teapicker/pkg/employee"
	"amberovsky/teapicker/pkg/member"
	"github.com/google/uuid"
)

/*
Service level for employees management
*/

type Service struct {
	repository       employee.Repository
	memberRepository member.Repository
}

func NewService(repository employee.Repository, memberRepository member.Repository) *Service {
	return &Service{
		repository:       repository,
		memberRepository: memberRepository,
	}
}

func (s *Service) GetAll() ([]*employee.Employee, error) {
	return s.repository.GetAll()
}

func (s *Service) Add(name string) error {
	return s.repository.Add(employee.NewEmployee(uuid.New(), name))
}

func (s *Service) Delete(uuidString string) error {
	employeeUUID := uuid.MustParse(uuidString)

	if err := s.repository.Delete(employeeUUID); err == nil {
		return s.memberRepository.DeleteEmployee(employeeUUID)
	} else {
		return err
	}
}
