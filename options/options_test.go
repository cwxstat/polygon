package options

import (
	"strings"
	"testing"
)

func TestOptions_List(t *testing.T) {

	tests := []struct {
		name           string
		setup          *Options
		ticker         string
		contractType   string
		expirationDate string
		want           string
		wantErr        bool
	}{
		{
			name:           "Simple",
			setup:          NewOptions(),
			ticker:         "CFLT",
			contractType:   "call",
			expirationDate: "2021-12-10",
			want:           "",
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := tt.setup
			o.Set(tt.ticker, tt.contractType, tt.expirationDate)

			got, err := o.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("Options.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.Contains(got, "OK") {
				t.Errorf("Options.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
