package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Post("/account", CreateAccount)
	app.Get("/account/:name", GetAccount)
	app.Patch("/account/:name", UpdateAmount)
	app.Get("/account", GetAllAccounts)
	app.Delete("/account/:name", DeleteAccount)
	app.Put("/account/:name", UpdateAccountName)

	log.Fatal(app.Listen(":8080"))
}
