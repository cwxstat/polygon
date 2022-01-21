package main

import (
	"fmt"

	"github.com/cwxstat/polygon/setup"
	"github.com/parnurzeal/gorequest"
)

func main() {

	key := setup.NewK().Key()
	ticker := "AAPL"

	url := fmt.Sprintf("https://api.polygon.io/v2/ticks/stocks/trades/%s/2020-10-14?reverse=true&limit=100&apiKey=%s", ticker, key)

	request := gorequest.New()
	resp, body, err := request.Get(url).End()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	_ = resp
	fmt.Println(body)

}
