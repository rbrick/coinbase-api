package coinbase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	PathCurrencies = "currencies"
	PathSpotPrice  = "prices/%s-%s/spot"

	TimePattern = "2006-01-02"
)

// Represents a list of accepted fiat currencies
// from endpoint: /currencies
type TraditionalCurrency struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	MinSize string `json:"min_size"`
}

// Represents the price of a specific coin.
// used for the prices endpoints
type Price struct {
	Amount   float64 `json:"amount,string"`
	Currency string  `json:"currency"`
}

func (c *Client) GetCurrencies(ctx context.Context) (currencies []TraditionalCurrency, err error) {
	req, err := c.makeRequest(http.MethodGet, PathCurrencies, url.Values{})

	if err != nil {
		return nil, err
	}

	err = c.execute(ctx, req, &currencies)
	return
}

func (c *Client) GetSpotPrice(ctx context.Context, crypto, fiat string, at time.Time) (price *Price, err error) {
	urlValues := url.Values{}
	dateQuery := at.Format(TimePattern)

	if !at.IsZero() {
		urlValues.Add("date", dateQuery)
	}

	url := fmt.Sprintf(PathSpotPrice, crypto, fiat)

	request, err := c.makeRequest(http.MethodGet, url, urlValues)

	if err != nil {
		return nil, err
	}

	price = &Price{}
	err = c.execute(ctx, request, price)
	return
}
