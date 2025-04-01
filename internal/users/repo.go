package users

type UserRepo struct {
	userModel *UserModel
}

func (userRepo *UserRepo) SetupRepo(userModel *UserModel) *UserRepo {
	userRepo.userModel = userModel
	return userRepo
}
