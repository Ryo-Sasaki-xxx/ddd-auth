package domains

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserEmail(t *testing.T) {

	t.Run("success NewUserEmail()", func(t *testing.T) {
		email := "hoge@fuga.com"
		userEmail, _ := NewUserEmail(email)

		assert.Equal(t, userEmail.email, email)
	})

	t.Run("validation error NewUserEmail()", func(t *testing.T) {
		email := "hogefuga.com"
		_, err := NewUserEmail(email)

		assert.Error(t, err)
	})
}
