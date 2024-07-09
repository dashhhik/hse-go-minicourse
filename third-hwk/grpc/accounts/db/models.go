package db

import "sync"

type Account struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type AccountStorage struct {
	mu      sync.RWMutex
	storage map[string]*Account
}

type ChangeAccountNameParams struct {
	NewName string `json:"new_name"`
}

type UpdateBalanceParams struct {
	Balance float64 `json:"balance"`
}

type UpdateNameParams struct {
	Name string `json:"name"`
}
