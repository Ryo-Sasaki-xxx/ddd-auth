package infra

import (
	"testing"

	domains "domains/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryUserRepository(t *testing.T) {

	t.Run("true Find()", func(t *testing.T) {
		inMemoryUserRepository := NewInMemoryUserRepository()
		exist := inMemoryUserRepository.Find(*domains.Inmemory_user_id)
		assert.True(t, exist)
	})
	t.Run("false Find()", func(t *testing.T) {
		inMemoryUserRepository := NewInMemoryUserRepository()
		uuid, _ := uuid.NewRandom()
		userId := domains.NewUserId(uuid)
		exist := inMemoryUserRepository.Find(*userId)
		assert.False(t, exist)
	})
	t.Run("true Save()", func(t *testing.T) {
		inMemoryUserRepository := NewInMemoryUserRepository()
		email := "hoge@fuga.com"
		userEmail, _ := domains.NewUserEmail(email)
		uuid, _ := uuid.NewRandom()
		userId := domains.NewUserId(uuid)
		name := "Ryo-Sasaki-xxx"
		userName, _ := domains.NewUserName(name)
		user, _ := domains.NewUser(*userEmail, *userId, *userName)

		err := inMemoryUserRepository.Save(*user)
		assert.Nil(t, err)
	})

}
