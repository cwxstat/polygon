package industry

import (
	"reflect"
	"testing"
)

func TestINDUSTRY_Industry(t *testing.T) {
	type fields struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		want    [][]string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Public test industry",
			fields:  fields{},
			want:    [][]string{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := NewIndustry()
			got, err := i.buildIndustry()
			if (err != nil) != tt.wantErr {
				t.Errorf("INDUSTRY.Industry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < 100 {
				t.Errorf("length no good")
				return
			}

		})
	}
}

func TestINDUSTRY_Tickers(t *testing.T) {
	type fields struct {
		url      string
		values   [][]string
		sector   map[string]int
		industry map[string]int
	}
	type args struct {
		industry string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name:   "Test tickers",
			fields: fields{},
			args:   args{},
			want:   []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := NewIndustry()
			if got := i.Tickers(tt.args.industry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("INDUSTRY.Tickers() = %v, want %v", got, tt.want)
			}
		})
	}
}
