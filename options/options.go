package options

import (
	"fmt"
	"github.com/cwxstat/polygon/setup"
	"github.com/parnurzeal/gorequest"
	"strings"
)

var (
	url = "https://api.polygon.io/v3/reference/options/contracts?underlying_ticker={ticker}&contract_type={contract_type}&expiration_date={expiration_date}&apiKey={apiKey}"
)

type Options struct {
	url string
	key string
}

func NewOptions() *Options {
	o := &Options{}
	o.key = setup.NewK().Key()
	return o
}

func (o *Options) Set(ticker, contract_type, expiration_date string) *Options {
	o.url = strings.Replace(url, "{apiKey}", o.key, 1)
	o.url = strings.Replace(o.url, "{ticker}", ticker, 1)
	o.url = strings.Replace(o.url, "{contract_type}", contract_type, 1)
	o.url = strings.Replace(o.url, "{expiration_date}", expiration_date, 1)
	return o

}

func (o *Options) List() (string, error) {
	request := gorequest.New()
	resp, body, err := request.Get(o.url).End()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return "", err[0]
	}
	_ = resp
	return body, nil

}
