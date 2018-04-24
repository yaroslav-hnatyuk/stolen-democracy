package entities

type UserRepository interface {
	Store(user User)
	FindById(id int) User
}
