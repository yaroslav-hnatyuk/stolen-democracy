package entities

type UserRepository interface {
	Store(user User)
	FindById(id int) User
}

type User struct {
	Id    int
	Name  string
	Login string
	Level int
}
