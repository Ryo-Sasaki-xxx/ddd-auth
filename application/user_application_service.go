package application

import (
	models "domains/models"
	service "domains/service"
)

type UserApplicationService struct {
	userRepository models.IUserRepository
	userService    service.UserService
}
