package usecases

import "github.com/yaroslav-hnatyuk/stolen-democracy/entities"

type Logger interface {
	Log(message string) error
}

type UserGetUsecase struct {
	UserId         int
	UserRepository entities.UserRepository
	Logger         Logger
}

func (uc *UserGetUsecase) execute() entities.Entity {
	user := uc.UserRepository.FindById(uc.UserId)
	if user == nil {
		uc.Logger.Log("User not found")
	}

	return user
}
