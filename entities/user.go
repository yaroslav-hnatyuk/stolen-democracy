package entities

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
	Level int    `json:"level"`
}
