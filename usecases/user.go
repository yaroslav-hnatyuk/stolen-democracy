package usecases

import "github.com/yaroslav-hnatyuk/stolen-democracy/entities"

type Logger interface {
	Log(message string) error
}

type Presenter interface {
	GetOutput(entities.User) []byte
}

type UserUsecase struct {
	UserId         int
	UserRepository entities.UserRepository
	Presenter      Presenter
}

func (uc UserUsecase) Execute() []byte {
	user := uc.UserRepository.FindById(uc.UserId)
	return uc.Presenter.GetOutput(user)
}
