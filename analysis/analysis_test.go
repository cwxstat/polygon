package analysis

import "testing"

func TestANALYSIS_Cov(t *testing.T) {
	type fields struct {
		a string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "",
			fields: fields{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAnalysis()
			a.Cov()
		})
	}
}
