package coinbase

import (
	"net/http"
	"time"
)

const (
	PathAccounts = "accounts"
)

const (
	WalletAccount = "wallet"
	FiatAccount   = "fiat"
	VaultAccount  = "vault"
)

type Account struct {
	Resource

	//Name is the Account name
	Name string `json:"name"`
	//Primary tells whether this account is the primary account
	Primary bool `json:"primary,omitempty"`
	//Type The type of the account. Can be "wallet", "fiat", or "vault"
	Type string `json:"type"`
	//Currency is the type of currency for this account
	Currency *Currency `json:"currency,omitempty"`
	//Balance is the current balance of the account
	Balance *Balance `json:"balance,omitempty"`

	//Timestamps for creation and last account action time
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Balance struct {
	Amount   float64 `json:"amount,string,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

type Accounts struct {
	Pagination
	Accounts []Account `data:"true"`
}

func (c *Client) Accounts(settings *Pagination) (*Accounts, error) {
	req, err := c.makeRequest(http.MethodGet, PathAccounts, settings.Encode())

	if err != nil {
		return nil, err
	}

	var accounts Accounts

	if err = c.execute(req, &accounts); err != nil {
		return nil, err
	}

	return &accounts, nil
}
