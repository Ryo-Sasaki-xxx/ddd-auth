package domains

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func TestUserId(t *testing.T) {

	t.Run("success NewUserId()", func(t *testing.T) {
		uuid, _ := uuid.NewRandom()

		userId := NewUserId(uuid)

		assert.Equal(t, userId.id, uuid.String())
	})
}
