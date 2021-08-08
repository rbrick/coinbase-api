package coinbase

import "time"

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
	Currency string `json:"currency,omitempty"`
	//Balance is the current balance of the account
	Balance *Balance `json:"balance,omitempty"`

	//Timestamps for creation and last account action time
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Balance struct {
	Amount   string `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
}
