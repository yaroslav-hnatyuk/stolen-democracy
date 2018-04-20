package main

import "github.com/yaroslav-hnatyuk/stolen-democracy/interfaces/controllers"

func main() {
	controller := controllers.UserController{}
	controller.GetUser(999)
}
