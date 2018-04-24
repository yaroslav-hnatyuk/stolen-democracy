package usecases

import (
	"github.com/yaroslav-hnatyuk/stolen-democracy/entities"
)

type UserGetUsecase struct {
	UserId         int
	UserRepository entities.UserRepository
	Presenter      PresenterInterface
}

func (uc UserGetUsecase) Execute() []byte {
	user := uc.UserRepository.FindById(uc.UserId)
	return uc.Presenter.GetOutput(user)
}
