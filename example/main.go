package main

import (
	"fmt"

	"github.com/rbrick/coinbase"
)

func main() {
	client := coinbase.New("", "", func(e error) {
		if e != nil {
			panic(e)
		}
	})

	fmt.Println(client.GetCurrencies())
}
