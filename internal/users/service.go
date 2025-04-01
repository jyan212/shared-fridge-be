package users

type UserService struct {
	userRepo *UserRepo
}

func (userService *UserService) SetupService(userRepo *UserRepo) *UserService {
	userService.userRepo = userRepo
	return userService
}
