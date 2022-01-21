package stocks

import (
	"reflect"
	"testing"
)

func TestStocks_MarketDays(t *testing.T) {

	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Simple first day",
			want:    "2021-01-04",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := NewStocks()
			got, err := o.buildMarketDays()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stocks.MarketDays() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got[0], tt.want) {
				t.Errorf("Stocks.MarketDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStocks_LastDay(t *testing.T) {

	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		{
			name:    "Simple length test",
			want:    10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := NewStocks()
			got, err := o.LastDay()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stocks.LastDay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("Stocks.LastDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStocks_HourlyValues(t *testing.T) {

	type args struct {
		ticker    string
		startDate string
		endDate   string
		limit     string
	}
	tests := []struct {
		name string

		args args
	}{
		{
			name: "",
			args: args{
				ticker:    "CFLT",
				startDate: "2021-12-16",
				endDate:   "2021-12-17",
				limit:     "5000",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := NewStocks()
			tickerData := o.HourlyValues(tt.args.ticker, tt.args.startDate, tt.args.endDate, tt.args.limit)
			if tickerData.Open()[14] != 63.7 {
				t.Errorf("tickerData.Open()[32] = %v, want %v", tickerData.Open()[14], 63.7)
			}
			tickerData = o.MinuteValues(tt.args.ticker, tt.args.startDate, tt.args.endDate, tt.args.limit)
			if tickerData.Open()[706] != 64.1 {
				t.Errorf("tickerData.Open()[32] = %v, want %v", tickerData.Open()[706], 64.1)
			}

		})
	}
}
