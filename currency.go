package coinbase

import (
	"net/http"
	"net/url"
)

const (
	PathCurrencies = "currencies"
)

// Represents a list of accepted fiat currencies
// from endpoint: /currencies
type Currency struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Size float64 `json:"min_size,string"`
}

func (c *Client) GetCurrencies() []Currency {
	req, err := c.makeRequest(http.MethodGet, PathCurrencies, url.Values{})

	c.errorHandler(err)

	var currencies []Currency

	c.errorHandler(c.execute(req, &currencies))

	return currencies
}
