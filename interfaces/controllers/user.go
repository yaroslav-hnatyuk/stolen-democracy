package controllers

import (
	"fmt"

	"github.com/yaroslav-hnatyuk/stolen-democracy/interfaces/presenters"
	"github.com/yaroslav-hnatyuk/stolen-democracy/interfaces/repositories"
	"github.com/yaroslav-hnatyuk/stolen-democracy/usecases"
)

type UserController struct{}

func (controller UserController) GetUser(userId int) {
	usecase := usecases.UserUsecase{
		UserId:         userId,
		UserRepository: repositories.UserRepository{},
		Presenter:      presenters.UserPresenter{},
	}

	output := usecase.Execute()
	fmt.Println(string(output))
}
