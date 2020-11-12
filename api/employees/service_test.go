package employees

import (
	"amberovsky/teapicker/pkg/employee"
	"amberovsky/teapicker/pkg/member"
	"github.com/google/uuid"
	"testing"
)

type EmployeeRepositoryMock struct {
	Employees []*employee.Employee
}

func (erm *EmployeeRepositoryMock) GetAll() ([]*employee.Employee, error) {
	return erm.Employees, nil
}

func (erm *EmployeeRepositoryMock) Add(e *employee.Employee) error {
	erm.Employees = append(erm.Employees, e)
	return nil
}

func (erm *EmployeeRepositoryMock) Delete(UUID uuid.UUID) error { return nil }

type MemberRepositoryMock struct {
	Members []*member.Member
}

func (mrm *MemberRepositoryMock) GetAllForTeam(teamUUID uuid.UUID) []*member.Member {
	return mrm.Members
}

func (mrm *MemberRepositoryMock) Get(UUID uuid.UUID) *member.Member { return mrm.Members[0] }

func (mrm *MemberRepositoryMock) Add(m *member.Member) error {
	mrm.Members = append(mrm.Members, m)
	return nil
}

func (mrm *MemberRepositoryMock) Delete(UUID uuid.UUID) error                 { return nil }
func (mrm *MemberRepositoryMock) DeleteAllFromTeam(teamUUID uuid.UUID) error  { return nil }
func (mrm *MemberRepositoryMock) DeleteEmployee(employeeUUID uuid.UUID) error { return nil }

func TestGetAll(t *testing.T) {
	service := NewService(&EmployeeRepositoryMock{
		Employees: []*employee.Employee{},
	}, &MemberRepositoryMock{
		Members: []*member.Member{},
	})

	employees, err := service.GetAll()

	if err != nil {
		t.Error(err)
	}

	if len(employees) != 0 {
		t.Errorf("expected non employees, got %d", len(employees))
	}
}

func TestAdd(t *testing.T) {
	service := NewService(&EmployeeRepositoryMock{
		Employees: []*employee.Employee{},
	}, &MemberRepositoryMock{
		Members: []*member.Member{},
	})

	if err := service.Add("name"); err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	service := NewService(&EmployeeRepositoryMock{
		Employees: []*employee.Employee{},
	}, &MemberRepositoryMock{
		Members: []*member.Member{},
	})

	if err := service.Delete(uuid.New().String()); err != nil {
		t.Error(err)
	}
}
