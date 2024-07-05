package db

import (
	"HSEGoCourse/second-hwk/accounts/models"
	"fmt"
	"sync"
)

var Accounts = NewAccountStorage()

type AccountStorage struct {
	mu      sync.RWMutex
	storage map[string]*models.Account
}

func NewAccountStorage() *AccountStorage {
	return &AccountStorage{
		storage: make(map[string]*models.Account),
	}
}

func (as *AccountStorage) CreateAccount(account *models.Account) {
	as.mu.Lock()
	defer as.mu.Unlock()
	as.storage[account.Name] = account
}

func (as *AccountStorage) GetAccount(name string) (*models.Account, error) {
	as.mu.RLock()
	defer as.mu.RUnlock()
	account, ok := as.storage[name]
	if !ok {
		return nil, fmt.Errorf("account with name %s not found", name)
	}
	return account, nil
}

func (as *AccountStorage) UpdateAmount(account *models.Account) error {
	as.mu.Lock()
	defer as.mu.Unlock()
	if _, ok := as.storage[account.Name]; ok {
		as.storage[account.Name] = account
		return nil
	} else {
		return fmt.Errorf("account with name %s not found", account.Name)
	}
}

func (as *AccountStorage) GetAllAccounts() []*models.Account {
	as.mu.RLock()
	defer as.mu.RUnlock()
	accounts := make([]*models.Account, 0, len(as.storage))
	for _, account := range as.storage {
		accounts = append(accounts, account)
	}
	return accounts
}
