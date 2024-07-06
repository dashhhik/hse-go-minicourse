package main

import (
	"HSEGoCourse/second-hwk/accounts/db"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func CreateAccount(c *fiber.Ctx) error {
	account := new(db.Account)
	if err := c.BodyParser(account); err != nil {
		log.Printf("Error parsing body: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	if account.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Account name is required"})
	}

	if account.Balance < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Account amount must be non-negative"})
	}

	db.Accounts.CreateAccount(account)

	return c.Status(fiber.StatusCreated).JSON(account)
}

func GetAccount(c *fiber.Ctx) error {
	name := c.Params("name")
	account, err := db.Accounts.GetAccount(name)
	if err != nil {
		log.Printf("Error fetching account: %v\n", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Account not found"})
	}

	return c.JSON(account)
}

func UpdateAmount(c *fiber.Ctx) error {
	name := c.Params("name")
	balance := new(db.UpdateBalanceParams)

	if err := c.BodyParser(balance); err != nil {
		log.Printf("Error parsing body: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	if balance.Balance < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Account amount must be non-negative"})
	}

	account := &db.Account{Name: name, Balance: balance.Balance}

	err := db.Accounts.UpdateAmount(account)
	if err != nil {
		log.Printf("Error updating account amount: %v\n", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Account not found"})
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetAllAccounts(c *fiber.Ctx) error {
	accounts := db.Accounts.GetAllAccounts()
	return c.JSON(accounts)
}

func DeleteAccount(c *fiber.Ctx) error {
	name := c.Params("name")

	err := db.Accounts.DeleteAccount(name)
	if err != nil {
		log.Printf("Error deleting account: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete account"})
	}

	return c.SendStatus(fiber.StatusOK)
}

func UpdateAccountName(c *fiber.Ctx) error {
	oldName := c.Params("name")
	newName := new(db.ChangeAccountNameParams)

	if err := c.BodyParser(newName); err != nil {
		log.Printf("Error parsing body: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	fmt.Println(oldName, newName.NewName)

	log.Printf("Received request to change account name from %s to %s\n", oldName, newName.NewName)

	err := db.Accounts.ChangeAccountName(newName.NewName, oldName)
	if err != nil {
		log.Printf("Error changing account name: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Printf("Account name changed successfully from %s to %s\n", oldName, newName.NewName)
	return c.SendStatus(fiber.StatusOK)
}
