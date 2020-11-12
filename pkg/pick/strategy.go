package pick

import "github.com/google/uuid"

type Strategy func(teamUUID uuid.UUID) (memberUUID *uuid.UUID, err error)
