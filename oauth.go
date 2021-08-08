package coinbase

import (
	"context"

	"golang.org/x/oauth2"
)

const (
	CoinbaseAuthorizeURL   = "https://www.coinbase.com/oauth/authorize"
	CoinbaseAccessTokenURL = "http://www.coinbase.com/oauth/token"
)

var (
	// The Endpoint for the OAuth2
	Endpoint = oauth2.Endpoint{
		AuthURL:  CoinbaseAuthorizeURL,
		TokenURL: CoinbaseAccessTokenURL,
	}
)

//NewOauthClient creates a new Coinbase client authenticated via OAuth2 token
func NewOauthClient(token oauth2.TokenSource) *Client {
	return &Client{
		httpClient: oauth2.NewClient(context.Background(), token),
	}
}
