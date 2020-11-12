package teams

import (
	"amberovsky/teapicker/pkg/member"
	"amberovsky/teapicker/pkg/team"
	"github.com/google/uuid"
	"testing"
)

type TeamRepositoryMock struct {
	Teams []*team.Team
}

func (trm *TeamRepositoryMock) GetAll() ([]*team.Team, error) {
	return trm.Teams, nil
}

func (trm *TeamRepositoryMock) Add(t *team.Team) error {
	trm.Teams = append(trm.Teams, t)
	return nil
}

func (trm *TeamRepositoryMock) Delete(UUID uuid.UUID) error { return nil }

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
	service := NewService(&TeamRepositoryMock{
		Teams: []*team.Team{},
	}, &MemberRepositoryMock{
		Members: []*member.Member{},
	})

	teams, err := service.GetAll()

	if err != nil {
		t.Error(err)
	}

	if len(teams) != 0 {
		t.Errorf("expected non teams, got %d", len(teams))
	}
}

func TestAdd(t *testing.T) {
	service := NewService(&TeamRepositoryMock{
		Teams: []*team.Team{},
	}, &MemberRepositoryMock{
		Members: []*member.Member{},
	})

	if err := service.Add("name"); err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	service := NewService(&TeamRepositoryMock{
		Teams: []*team.Team{},
	}, &MemberRepositoryMock{
		Members: []*member.Member{},
	})

	if err := service.Delete(uuid.New().String()); err != nil {
		t.Error(err)
	}
}
