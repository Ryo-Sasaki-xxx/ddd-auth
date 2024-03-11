package domains

import (
	"testing"

	domains "domains/models"
	infra "infra/in_memory"

	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {

	t.Run("true Exist()", func(t *testing.T) {
		inMemoryUserRepository := infra.NewInMemoryUserRepository()
		userService := NewUserService(inMemoryUserRepository)
		email := "hoge@fuga.com"
		userEmail, _ := domains.NewUserEmail(email)
		name := "Ryo-Sasaki-xxx"
		userName, _ := domains.NewUserName(name)

		user, _ := domains.NewUser(*userEmail, *domains.Inmemory_user_id, *userName)
		assert.True(t, userService.Exist(*user))
	})

	// t.Run("false Exist()", func(t *testing.T) {
	// 	userService := NewUserService()
	// 	email := "hoge@fuga.com"
	// 	userEmail, _ := domains.NewUserEmail(email)
	// 	uuid, _ := uuid.NewRandom()
	// 	userId := domains.NewUserId(uuid)
	// 	name := "Ryo-Sasaki-xxx"
	// 	userName, _ := domains.NewUserName(name)

	// 	user, _ := domains.NewUser(*userEmail, *userId, *userName)
	// 	assert.False(t, userService.Exist(*user))
	// })
}
