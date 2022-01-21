package stocks

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/cwxstat/polygon/request"
	"github.com/cwxstat/polygon/setup"
)

var (
	ticker          = "AAPL" // Some standard stock, only for date maket open
	startDate       = "2021-01-01"
	endDate         = "2041-07-22" // Way in the future
	url             = "https://api.polygon.io/v2/last/trade/{ticker}?apiKey={apiKey}"
	urlEntireMarket = "https://api.polygon.io/v2/aggs/grouped/locale/us/market/stocks/{date}?adjusted=true&apiKey={apiKey}"
	urlDays         = fmt.Sprintf("https://api.polygon.io/v2/aggs/ticker/%s/range/1/day/%s/%s?adjusted=true&sort=asc&limit=50000&apiKey={apiKey}",
		ticker, startDate, endDate)
	urlHour   = "https://api.polygon.io/v2/aggs/ticker/{ticker}/range/1/hour/{startDate}/{endDate}?adjusted=true&sort=asc&limit={limit}&apiKey={apiKey}"
	urlMinute = "https://api.polygon.io/v2/aggs/ticker/{ticker}/range/1/minute/{startDate}/{endDate}?adjusted=true&sort=asc&limit={limit}&apiKey={apiKey}"
)

// https://mholt.github.io/json-to-go/
type AggregateJSON struct {
	Ticker       string `json:"ticker"`
	QueryCount   int    `json:"queryCount"`
	ResultsCount int    `json:"resultsCount"`
	Adjusted     bool   `json:"adjusted"`
	Results      []struct {
		V  int     `json:"v"`
		Vw float64 `json:"vw"`
		O  float64 `json:"o"`
		C  float64 `json:"c"`
		H  float64 `json:"h"`
		L  float64 `json:"l"`
		T  int64   `json:"t"`
		N  int     `json:"n"`
	} `json:"results"`
	Status    string `json:"status"`
	RequestID string `json:"request_id"`
	Count     int    `json:"count"`
}

type EntireMarket []struct {
	Q  string  `json:"T"`
	V  int     `json:"v"`
	Vw float64 `json:"vw"`
	O  float64 `json:"o"`
	C  float64 `json:"c"`
	H  float64 `json:"h"`
	L  float64 `json:"l"`
	T  int64   `json:"t"`
	N  int     `json:"n"`
}

type Stocks struct {
	url             string
	urlEntireMarket string
	key             string
	marketDays      []string
}

func NewStocks() *Stocks {
	o := &Stocks{}
	o.key = setup.NewK().Key()
	o.urlEntireMarket = strings.Replace(urlEntireMarket, "{apiKey}", setup.NewK().Key(), 1)
	urlDays = strings.Replace(urlDays, "{apiKey}", setup.NewK().Key(), 1)
	o.url = strings.Replace(url, "{apiKey}", setup.NewK().Key(), 1)
	return o
}

type TickerData struct {
	timeStamp []string
	ticker    string
	o         []float64
	h         []float64
	l         []float64
	c         []float64
	v         []float64
	vw        []float64
}

func (t *TickerData) Ticker() string {
	return t.ticker
}

func (t *TickerData) Open() []float64 {
	return t.o
}

func (o *Stocks) dataByTimeSpan(urlInput, ticker, startDate, endDate, limit string) *TickerData {
	url := strings.Replace(urlInput, "{apiKey}", setup.NewK().Key(), 1)
	url = strings.Replace(url, "{ticker}", ticker, 1)
	url = strings.Replace(url, "{startDate}", startDate, 1)
	url = strings.Replace(url, "{endDate}", endDate, 1)
	url = strings.Replace(url, "{limit}", limit, 1)

	aggregateJSON := &AggregateJSON{}
	body, err := request.Request(url)
	if err != nil {
		return nil
	}
	json.Unmarshal([]byte(body), aggregateJSON)

	tic := &TickerData{}
	tic.timeStamp = []string{}
	tic.ticker = ticker

	tic.o = []float64{}
	tic.c = []float64{}
	tic.h = []float64{}
	tic.l = []float64{}
	tic.v = []float64{}
	tic.vw = []float64{}
	for _, v := range aggregateJSON.Results {
		t := time.Unix(0, v.T*int64(time.Millisecond))
		tic.timeStamp = append(tic.timeStamp, t.Format("2006-01-02 15:04:05"))

		tic.o = append(tic.o, v.O)
		tic.c = append(tic.c, v.C)
		tic.h = append(tic.h, v.H)
		tic.l = append(tic.l, v.L)
		tic.v = append(tic.v, float64(v.V))
		tic.vw = append(tic.vw, v.Vw)
	}

	return tic

}

func (o *Stocks) HourlyValues(ticker, startDate, endDate, limit string) *TickerData {
	return o.dataByTimeSpan(urlHour, ticker, startDate, endDate, limit)
}

func (o *Stocks) MinuteValues(ticker, startDate, endDate, limit string) *TickerData {
	return o.dataByTimeSpan(urlMinute, ticker, startDate, endDate, limit)
}

func (o *Stocks) set(ticker string) *Stocks {
	o.url = strings.Replace(url, "{apiKey}", o.key, 1)
	o.url = strings.Replace(o.url, "{ticker}", ticker, 1)
	return o

}

func (o *Stocks) buildMarketDays() ([]string, error) {

	aggregateJSON := &AggregateJSON{}
	body, err := request.Request(urlDays)
	if err != nil {
		return []string{}, nil
	}
	json.Unmarshal([]byte(body), aggregateJSON)

	r := []string{}
	for _, v := range aggregateJSON.Results {
		t := time.Unix(0, v.T*int64(time.Millisecond))
		r = append(r, t.Format("2006-01-02"))
	}
	o.marketDays = r
	return r, nil
}

func (o *Stocks) LastDay() (string, error) {

	if o.marketDays == nil {
		_, err := o.buildMarketDays()
		if err != nil {
			return "", err
		}
	}

	l := len(o.marketDays)
	return o.marketDays[l-1], nil
}
