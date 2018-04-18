package entities

type Entity interface {
	serialize()
}

type UserRepository interface {
	Store(user *User)
	FindById(id int) *User
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
	Level int    `json:"level"`
}

func (u User) serialize() string {
	return ""
}
