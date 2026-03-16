package manager

import (
	"fmt"
)

// Account represents a ChatGPT API account.
type Account struct {
	APIKey   string
	Quota    int
	Used     int
	Rotation bool
}

// AccountManager manages multiple ChatGPT API accounts.
type AccountManager struct {
	Accounts []Account
}

// NewAccountManager creates a new AccountManager.
func NewAccountManager() *AccountManager {
	return &AccountManager{Accounts: []Account{}}
}

// AddAccount adds a new account to the manager.
func (am *AccountManager) AddAccount(apiKey string, quota int) {
	am.Accounts = append(am.Accounts, Account{APIKey: apiKey, Quota: quota, Used: 0, Rotation: false})
}

// TrackUsage tracks the usage of a specific account.
func (am *AccountManager) TrackUsage(apiKey string, usage int) error {
	for i, account := range am.Accounts {
		if account.APIKey == apiKey {
			account.Used += usage
			if account.Used > account.Quota {
				return fmt.Errorf("Exceeded quota for account: %s", apiKey)
			}
			am.Accounts[i] = account
			return nil
		}
	}
	return fmt.Errorf("Account not found: %s", apiKey)
}

// RotateAccount returns the next available account for use.
func (am *AccountManager) RotateAccount() (*Account, error) {
	for i := range am.Accounts {
		if am.Accounts[i].Used < am.Accounts[i].Quota {
			am.Accounts[i].Rotation = true
			return &am.Accounts[i], nil
		}
	}
	return nil, fmt.Errorf("No accounts available for rotation")
}