package coinbase

import (
	"context"
	"net/http"
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
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Primary          bool     `json:"primary"`
	Type             string   `json:"type"`
	Currency         Currency `json:"currency"`
	Balance          Balance  `json:"balance"`
	CreatedAt        string   `json:"created_at"`
	UpdatedAt        string   `json:"updated_at"`
	AllowDeposits    bool     `json:"allow_deposits"`
	AllowWithdrawals bool     `json:"allow_withdrawals"`
}

type Currency struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	SortIndex    int64  `json:"sort_index"`
	Exponent     int64  `json:"exponent"`
	Type         string `json:"type"`
	AddressRegex string `json:"address_regex"`
	AssetID      string `json:"asset_id"`
	Slug         string `json:"slug"`
}

type Balance struct {
	Amount   float64 `json:"amount,string,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

type Accounts struct {
	Pagination
	Accounts []Account `data:"true"`
}

func (c *Client) Accounts(ctx context.Context, settings *Pagination) (*Accounts, error) {
	req, err := c.makeRequest(http.MethodGet, PathAccounts, settings.Encode())

	if err != nil {
		return nil, err
	}

	var accounts Accounts

	if err = c.execute(ctx, req, &accounts); err != nil {
		return nil, err
	}

	return &accounts, nil
}
