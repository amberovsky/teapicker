package members

import (
	"amberovsky/teapicker/pkg/member"
	"github.com/google/uuid"
	"testing"
)

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

func TestGetTeamMembersForExistingTeam(t *testing.T) {
	teamUUID := uuid.New()
	memberUUID := uuid.New()
	employeeUUID := uuid.New()

	service := NewService(&MemberRepositoryMock{
		Members: []*member.Member{
			member.NewMember(memberUUID, teamUUID, employeeUUID),
		},
	})

	teamMembers := service.GetTeamMembers(teamUUID)
	if len(teamMembers) != 1 {
		t.Errorf("expected only 1 team member, got %d", len(teamMembers))
	}

	if teamMembers[0].UUID.String() != memberUUID.String() {
		t.Errorf("received different UUIDs: %s vs %s", teamMembers[0].UUID.String(), memberUUID.String())
	}
}

func TestGetTeamMembersForNonExistingTeam(t *testing.T) {
	service := NewService(&MemberRepositoryMock{
		Members: []*member.Member{},
	})

	teamMembers := service.GetTeamMembers(uuid.New())
	if len(teamMembers) != 0 {
		t.Errorf("expected none team member, got %d", len(teamMembers))
	}
}

func TestGetExistingTeamMember(t *testing.T) {
	teamUUID := uuid.New()
	memberUUID := uuid.New()
	employeeUUID := uuid.New()

	service := NewService(&MemberRepositoryMock{
		Members: []*member.Member{
			member.NewMember(memberUUID, teamUUID, employeeUUID),
		},
	})

	teamMember := service.Get(memberUUID)
	if teamMember == nil {
		t.Errorf("expected 1 team team member, got nothing")
	} else {
		if teamMember.UUID.String() != memberUUID.String() {
			t.Errorf("received different UUIDs: %s vs %s", teamMember.UUID.String(), memberUUID.String())
		}
	}
}

func TestGetNonExistingTeamMember(t *testing.T) {
	service := NewService(&MemberRepositoryMock{
		Members: []*member.Member{
			nil,
		},
	})

	teamMember := service.Get(uuid.New())
	if teamMember != nil {
		t.Errorf("expected none team team member, got something")
	}
}

func TestAdd(t *testing.T) {
	service := NewService(&MemberRepositoryMock{
		Members: []*member.Member{},
	})

	if err := service.Add(uuid.New().String(), uuid.New().String()); err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	service := NewService(&MemberRepositoryMock{
		Members: []*member.Member{},
	})

	if err := service.Delete(uuid.New().String()); err != nil {
		t.Error(err)
	}
}
