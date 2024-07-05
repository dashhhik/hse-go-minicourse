package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Post("/account/create", CreateAccount)
	app.Get("/account/:name", GetAccount)
	app.Post("/account/update", UpdateAmount)
	app.Get("/account", GetAllAccounts)

	app.Listen(":8080")
}
