package repositories

import (
	"fmt"

	"github.com/yaroslav-hnatyuk/stolen-democracy/entities"
)

type UserRepository struct{}

func (ur UserRepository) Store(user entities.User) {
	fmt.Println("Save user to DB")
}

func (ur UserRepository) FindById(id int) entities.User {
	return entities.User{
		ID:    999,
		Name:  "John Smoth",
		Login: "js",
		Level: 85,
	}
}
