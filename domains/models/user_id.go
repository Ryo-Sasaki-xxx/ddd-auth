package domains

import (
	"github.com/google/uuid"
)

type UserId struct {
	id string
}

func NewUserId(id uuid.UUID) *UserId {
	userId := &UserId{}
	userId.id = id.String()

	return userId
}
