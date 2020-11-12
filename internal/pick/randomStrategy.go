package pick

import (
	"amberovsky/teapicker/pkg/member"
	"github.com/google/uuid"
	"math/rand"
)

type randomStrategy struct {
	memberRepository member.Repository
}

/*
In-memory storage implementation for picking the next tea maker
*/

func NewRandomStrategy(memberRepository member.Repository) *randomStrategy {
	return &randomStrategy{
		memberRepository: memberRepository,
	}
}

func (rs *randomStrategy) Invoke(teamUUID uuid.UUID) (memberUUID *uuid.UUID, err error) {
	members := rs.memberRepository.GetAllForTeam(teamUUID)
	if len(members) > 0 {
		return &members[rand.Intn(len(members))].UUID, nil
	} else {
		return nil, nil
	}
}
