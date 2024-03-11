package domains

import (
	domains "domains/models"
)

type UserService struct {
	userRepository domains.IUserRepository
}

func NewUserService(userRepository domains.IUserRepository) UserService {
	userService := UserService{}
	userService.userRepository = userRepository
	return userService
}

func (u UserService) Exist(user domains.User) bool {

	return u.userRepository.Find(user.Id)
}
