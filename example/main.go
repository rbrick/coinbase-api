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
	client := coinbase.New(apiKey, apiSecret, func(e error) {
		if e != nil {
			panic(e)
		}
	})

	fmt.Println("current BTC Price in USD is", client.GetSpotPrice("BTC", "USD", time.Time{}).Amount)

	past := time.Date(2018, time.August, 1, 0, 0, 0, 0, time.UTC)

	fmt.Println("price of BTC on Aug 1st 2018", client.GetSpotPrice("BTC", "USD", past).Amount)

	user := client.CurrentUser()

	fmt.Println(user)
}
