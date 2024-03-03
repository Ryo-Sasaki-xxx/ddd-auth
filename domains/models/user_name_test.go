package domains

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserName(t *testing.T) {

	t.Run("success NewUserName()", func(t *testing.T) {
		name := "Ryo-Sasaki-xxx"
		userName, _ := NewUserName(name)

		assert.Equal(t, userName.name, name)
	})

	t.Run("validation error NewUserName()", func(t *testing.T) {
		name := ""
		_, err := NewUserName(name)

		assert.Error(t, err)
	})
}
