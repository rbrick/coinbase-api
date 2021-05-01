package main

import (
	"fmt"
	"time"

	"github.com/rbrick/coinbase"
)

func main() {
	client := coinbase.New("", "", func(e error) {
		if e != nil {
			panic(e)
		}
	})

	fmt.Println("current BTC Price in USD is", client.GetSpotPrice("BTC", "USD", time.Time{}).Amount)

	past := time.Date(2018, time.August, 1, 0, 0, 0, 0, time.UTC)

	fmt.Println("price of BTC on Aug 1st 2018", client.GetSpotPrice("BTC", "USD", past).Amount)
}
