package domains

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {

	t.Run("success NewUser()", func(t *testing.T) {
		email := "hoge@fuga.com"
		userEmail, _ := NewUserEmail(email)
		uuid, _ := uuid.NewRandom()
		userId := NewUserId(uuid)
		name := "Ryo-Sasaki-xxx"
		userName, _ := NewUserName(name)

		user, _ := NewUser(*userEmail, *userId, *userName)

		assert.Equal(t, user.email, *userEmail)
		assert.Equal(t, user.id, *userId)
		assert.Equal(t, user.name, *userName)
	})
}
