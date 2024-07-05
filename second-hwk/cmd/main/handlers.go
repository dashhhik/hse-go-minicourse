package main

import (
	"HSEGoCourse/second-hwk/accounts/db"
	"HSEGoCourse/second-hwk/accounts/models"
	"github.com/gofiber/fiber/v2"
	"sync"
)

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func CreateAccount(c *fiber.Ctx) error {

	account := new(models.Account)
	if err := c.BodyParser(account); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if account.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "name is empty"})
	}

	if account.Amount < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "amount is negative"})
	}

	db.Accounts.CreateAccount(account)

	return c.SendStatus(fiber.StatusCreated)
}

func GetAccount(c *fiber.Ctx) error {
	name := c.Params("name")
	account, err := db.Accounts.GetAccount(name)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(account)
}

func UpdateAmount(c *fiber.Ctx) error {
	account := new(models.Account)
	if err := c.BodyParser(account); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if account.Amount < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "amount is negative"})
	}

	err := db.Accounts.UpdateAmount(account)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetAllAccounts(c *fiber.Ctx) error {
	accounts := db.Accounts.GetAllAccounts()
	return c.JSON(accounts)

}
