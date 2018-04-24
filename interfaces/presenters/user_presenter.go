package presenters

import (
	"encoding/json"

	"github.com/yaroslav-hnatyuk/stolen-democracy/entities"
	"github.com/yaroslav-hnatyuk/stolen-democracy/interfaces"
)

type UserPresenter struct {
	Logger interfaces.LoggerInterface
}

func (up UserPresenter) GetOutput(entity interface{}) []byte {
	user, ok := entity.(entities.User)
	if !ok {
		up.Logger.Log("Unable to present User. Expected User entity.")
	}

	result, err := json.Marshal(&user)
	if err != nil {
		up.Logger.Log("Unebale to present User. Unexpected problem with json.Marshal.")
	}

	return result

}
