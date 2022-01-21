package setup

import "testing"

func TestKEY_Key(t *testing.T) {

	tests := []struct {
		name  string
		start *KEY
		want  string
	}{
		{
			name:  "Simple",
			start: NewK(),
			want:  "Some valid key",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := tt.start
			if got := k.Key(); got == "" {
				t.Errorf("KEY.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}
