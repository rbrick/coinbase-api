package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/rbrick/coinbase"
)

var apiKey, apiSecret string

func init() {
	d, _ := ioutil.ReadFile("secrets")

	var m map[string]string

	json.Unmarshal(d, &m)

	apiKey = m["key"]
	apiSecret = m["secret"]
}

func main() {
	client := coinbase.New(apiKey, apiSecret)

	price, _ := client.GetSpotPrice("BTC", "USD", time.Time{})

	fmt.Println("current BTC Price in USD is", price.Amount)

	past := time.Date(2018, time.August, 1, 0, 0, 0, 0, time.UTC)

	price, _ = client.GetSpotPrice("BTC", "USD", past)
	fmt.Println("price of BTC on Aug 1st 2018", price.Amount)

	user, _ := client.CurrentUser()

	fmt.Println(user)

	pagination := &coinbase.Pagination{
		Limit: 25,
		Order: coinbase.Ascending,
	}

	fmt.Println(pagination.Encode().Encode())
}
