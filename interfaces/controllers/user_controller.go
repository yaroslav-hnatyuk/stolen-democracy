package controllers

import (
	"fmt"

	"github.com/yaroslav-hnatyuk/stolen-democracy/interfaces"
	"github.com/yaroslav-hnatyuk/stolen-democracy/interfaces/presenters"
	"github.com/yaroslav-hnatyuk/stolen-democracy/interfaces/repositories"
	"github.com/yaroslav-hnatyuk/stolen-democracy/usecases"
)

type UserController struct {
	Logger interfaces.LoggerInterface
}

func (controller UserController) GetUser(userId int) {
	presenter := presenters.UserPresenter{
		Logger: controller.Logger,
	}

	usecase := usecases.UserGetUsecase{
		UserId:         userId,
		UserRepository: repositories.UserRepository{},
		Presenter:      presenter,
	}

	output := usecase.Execute()
	fmt.Println(string(output))
}
