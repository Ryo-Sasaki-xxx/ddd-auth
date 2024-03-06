package infra

import domains "github.com/Ryo-Sasaki-xxx/ddd-auth/domains/models"

type InMemoryUserRepository struct {
	store map[domains.UserId]*domains.User
}

func NewInMemoryUserRepository() InMemoryUserRepository {
	inMemoryUserRepository := InMemoryUserRepository{}
	store := make(map[domains.UserId]*domains.User)
	email := "hoge@fuga.com"
	userEmail, _ := domains.NewUserEmail(email)
	name := "Ryo-Sasaki-xxx"
	userName, _ := domains.NewUserName(name)
	exist_user, _ := domains.NewUser(*userEmail, *domains.Inmemory_user_id, *userName)
	store[*domains.Inmemory_user_id] = exist_user
	inMemoryUserRepository.store = store

	return inMemoryUserRepository
}

func (u InMemoryUserRepository) Find(id domains.UserId) bool {
	user := u.store[id]

	return user != nil
}
