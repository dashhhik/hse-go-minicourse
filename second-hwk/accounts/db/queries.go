package db

import (
	"fmt"
)

func NewAccountStorage() *AccountStorage {
	return &AccountStorage{
		storage: make(map[string]*Account),
	}
}

func (as *AccountStorage) CreateAccount(account *Account) {
	as.mu.Lock()
	defer as.mu.Unlock()
	as.storage[account.Name] = account
}

func (as *AccountStorage) GetAccount(name string) (*Account, error) {
	as.mu.RLock()
	defer as.mu.RUnlock()
	account, ok := as.storage[name]
	if !ok {
		return nil, fmt.Errorf("account with name %s not found", name)
	}
	return account, nil
}

func (as *AccountStorage) UpdateAmount(account *Account) error {
	as.mu.Lock()
	defer as.mu.Unlock()
	storedAccount, ok := as.storage[account.Name]
	if !ok {
		return fmt.Errorf("account with name %s not found", account.Name)
	}
	storedAccount.Balance = account.Balance
	return nil
}

func (as *AccountStorage) GetAllAccounts() []*Account {
	as.mu.RLock()
	defer as.mu.RUnlock()
	accounts := make([]*Account, 0, len(as.storage))
	for _, account := range as.storage {
		accounts = append(accounts, account)
	}
	return accounts
}

func (as *AccountStorage) DeleteAccount(name string) error {
	as.mu.Lock()
	defer as.mu.Unlock()
	if _, ok := as.storage[name]; !ok {
		return fmt.Errorf("account with name %s not found", name)
	}
	delete(as.storage, name)
	return nil
}

func (as *AccountStorage) ChangeAccountName(newName, oldName string) error {
	as.mu.Lock()
	defer as.mu.Unlock()

	account, ok := as.storage[oldName]
	if !ok {
		return fmt.Errorf("account with name %s not found", oldName)
	}

	if _, exists := as.storage[newName]; exists {
		return fmt.Errorf("account with name %s already exists", newName)
	}

	delete(as.storage, oldName)

	account.Name = newName
	as.storage[newName] = account

	return nil
}
