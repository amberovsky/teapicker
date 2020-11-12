package member

import (
	"amberovsky/teapicker/pkg/member"
	"github.com/google/uuid"
)

/*
In-memory storage implementation for teams composition
*/

type inMemoryRepository struct {
	members map[string]*member.Member
}

func NewInMemoryRepository() *inMemoryRepository {
	members := make(map[string]*member.Member)

	// Been to Rome
	members["cffe92a3-1bbc-4614-839f-974d030e805d"] = member.NewMember(uuid.MustParse("cffe92a3-1bbc-4614-839f-974d030e805d"), uuid.MustParse("549a825d-78bd-4d30-9571-d89532cc6932"), uuid.MustParse("0a2206ef-4f84-48a8-a7fc-18bae866a488"))
	members["5f630c5e-b7d7-464c-868e-bb5a8c089ce8"] = member.NewMember(uuid.MustParse("5f630c5e-b7d7-464c-868e-bb5a8c089ce8"), uuid.MustParse("549a825d-78bd-4d30-9571-d89532cc6932"), uuid.MustParse("b6625257-19eb-4f98-a05e-c0674db6d870"))
	members["54aecaf3-e7e2-45a4-8cc4-abfe77ad96b4"] = member.NewMember(uuid.MustParse("54aecaf3-e7e2-45a4-8cc4-abfe77ad96b4"), uuid.MustParse("549a825d-78bd-4d30-9571-d89532cc6932"), uuid.MustParse("63c5646c-e212-4d14-98e4-1b0979a1314f"))
	members["93b13b4a-225c-4deb-b9cf-0164085d62aa"] = member.NewMember(uuid.MustParse("93b13b4a-225c-4deb-b9cf-0164085d62aa"), uuid.MustParse("549a825d-78bd-4d30-9571-d89532cc6932"), uuid.MustParse("872d921f-ad5a-4354-8564-4657178dda01"))
	members["f4b9afdc-7af5-4f8c-be87-411954e7cf07"] = member.NewMember(uuid.MustParse("f4b9afdc-7af5-4f8c-be87-411954e7cf07"), uuid.MustParse("549a825d-78bd-4d30-9571-d89532cc6932"), uuid.MustParse("b7977c1d-7528-4a12-8e47-044840ae5988"))
	members["a4e12c27-cbcb-48a3-897f-919ae2780809"] = member.NewMember(uuid.MustParse("a4e12c27-cbcb-48a3-897f-919ae2780809"), uuid.MustParse("549a825d-78bd-4d30-9571-d89532cc6932"), uuid.MustParse("9d0f24ac-f7c0-40b8-97e4-b3bdcb1bd49e"))
	members["243654c2-2275-4608-8152-880985871fc1"] = member.NewMember(uuid.MustParse("243654c2-2275-4608-8152-880985871fc1"), uuid.MustParse("549a825d-78bd-4d30-9571-d89532cc6932"), uuid.MustParse("a96dd9f6-8e45-4b69-ac32-a153f590c62c"))
	members["c651c6b8-5357-43af-bbb9-29c29d6eedf8"] = member.NewMember(uuid.MustParse("c651c6b8-5357-43af-bbb9-29c29d6eedf8"), uuid.MustParse("549a825d-78bd-4d30-9571-d89532cc6932"), uuid.MustParse("8a86fa37-64c7-461a-b860-852041a1b9a5"))
	members["1b88df69-3df3-4d1b-aae1-f77a2153059c"] = member.NewMember(uuid.MustParse("1b88df69-3df3-4d1b-aae1-f77a2153059c"), uuid.MustParse("549a825d-78bd-4d30-9571-d89532cc6932"), uuid.MustParse("216551a6-11d1-4d30-89fd-983fc8daf345"))

	// Founders
	members["0c9a17a4-1e66-42dc-a197-ce62d3b8fbfb"] = member.NewMember(uuid.MustParse("0c9a17a4-1e66-42dc-a197-ce62d3b8fbfb"), uuid.MustParse("b8480227-61c2-44d2-8c14-838db674cc39"), uuid.MustParse("5167fa62-72b3-4397-b69a-2c0bef9391d4"))
	members["db6cbf01-aabe-4fba-966a-4ba7afaf67c9"] = member.NewMember(uuid.MustParse("db6cbf01-aabe-4fba-966a-4ba7afaf67c9"), uuid.MustParse("b8480227-61c2-44d2-8c14-838db674cc39"), uuid.MustParse("eb3973c6-ed75-4a62-b9e8-4ba7226b5469"))
	members["7eb8827b-8516-4a2e-a632-0a7cf65254f1"] = member.NewMember(uuid.MustParse("7eb8827b-8516-4a2e-a632-0a7cf65254f1"), uuid.MustParse("b8480227-61c2-44d2-8c14-838db674cc39"), uuid.MustParse("659a3d79-3ec3-42cc-8540-310a2599ad63"))

	return &inMemoryRepository{members: members}
}

func (r *inMemoryRepository) GetAllForTeam(teamUUID uuid.UUID) []*member.Member {
	members := []*member.Member{}
	for _, value := range r.members {
		if value.TeamUUID.String() == teamUUID.String() {
			members = append(members, value)
		}
	}

	return members
}

func (r *inMemoryRepository) Get(UUID uuid.UUID) *member.Member {
	return r.members[UUID.String()]
}

func (r *inMemoryRepository) Add(m *member.Member) error {
	r.members[m.UUID.String()] = m
	return nil
}

func (r *inMemoryRepository) Delete(UUID uuid.UUID) error {
	delete(r.members, UUID.String())
	return nil
}

func (r *inMemoryRepository) DeleteAllFromTeam(teamUUID uuid.UUID) error {
	for key, value := range r.members {
		if value.TeamUUID.String() == teamUUID.String() {
			delete(r.members, key)
		}
	}

	return nil
}

func (r *inMemoryRepository) DeleteEmployee(employeeUUID uuid.UUID) error {
	for key, value := range r.members {
		if value.EmployeeUUID.String() == employeeUUID.String() {
			delete(r.members, key)
		}
	}

	return nil
}
