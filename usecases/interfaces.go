package usecases

type PresenterInterface interface {
	GetOutput(entity interface{}) []byte
}
