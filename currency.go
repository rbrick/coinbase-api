package coinbase

import (
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
type Currency struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Size float64 `json:"min_size,string"`
}

// Represents the price of a specific coin.
// used for the prices endpoints
type Price struct {
	Amount   float64 `json:"amount,string"`
	Currency string  `json:"currency"`
}

func (c *Client) GetCurrencies() []Currency {
	var currencies []Currency

	req, err := c.makeRequest(http.MethodGet, PathCurrencies, url.Values{})

	c.errorHandler(err)

	c.errorHandler(c.execute(req, &currencies))

	return currencies
}

func (c *Client) GetSpotPrice(crypto, fiat string, at time.Time) *Price {
	var price Price

	urlValues := url.Values{}
	dateQuery := at.Format(TimePattern)

	if !at.IsZero() {
		urlValues.Add("date", dateQuery)
	}

	url := fmt.Sprintf(PathSpotPrice, crypto, fiat)

	request, err := c.makeRequest(http.MethodGet, url, urlValues)

	c.errorHandler(err)
	c.errorHandler(c.execute(request, &price))

	return &price
}
