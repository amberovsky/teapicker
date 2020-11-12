package team

import (
	"amberovsky/teapicker/pkg/team"
	"github.com/google/uuid"
)

type inMemoryRepository struct {
	teams map[string]*team.Team
}

/*
In-memory storage implementation for teams
*/

func NewInMemoryRepository() *inMemoryRepository {
	teams := make(map[string]*team.Team)

	teams["549a825d-78bd-4d30-9571-d89532cc6932"] = team.NewTeam(uuid.MustParse("549a825d-78bd-4d30-9571-d89532cc6932"), "Been to Rome")
	teams["b8480227-61c2-44d2-8c14-838db674cc39"] = team.NewTeam(uuid.MustParse("b8480227-61c2-44d2-8c14-838db674cc39"), "Founders")

	return &inMemoryRepository{teams: teams}
}

func (r *inMemoryRepository) GetAll() ([]*team.Team, error) {
	teams := []*team.Team{}
	for _, value := range r.teams {
		teams = append(teams, value)
	}

	return teams, nil
}

func (r *inMemoryRepository) Add(t *team.Team) error {
	r.teams[t.UUID.String()] = t
	return nil
}

func (r *inMemoryRepository) Delete(UUID uuid.UUID) error {
	delete(r.teams, UUID.String())
	return nil
}
