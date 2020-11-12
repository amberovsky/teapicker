package pick

import (
	"github.com/google/uuid"
	"testing"
)

func TestDeleteCallsRepository(t *testing.T) {
	UUID := uuid.New()

	service := NewService(func(teamUUID uuid.UUID) (memberUUID *uuid.UUID, err error) {
		return &UUID, nil
	})

	picked, err := service.Pick(UUID.String())
	if err != nil {
		t.Error(err)
	}

	if picked.String() != UUID.String() {
		t.Errorf("received different UUIDs: %s vs %s", UUID.String(), picked.String())
	}
}
