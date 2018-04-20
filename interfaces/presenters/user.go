package presenters

import (
	"encoding/json"

	"github.com/yaroslav-hnatyuk/stolen-democracy/entities"
)

type UserPresenter struct{}

func (up UserPresenter) GetOutput(user entities.User) []byte {
	res, _ := json.Marshal(&user)
	return res
}
