package bitbu

import "testing"

func TestDefaultDataBit_BitType(t *testing.T) {
	tests := []struct {
		name string
		b    DefaultDataBit
		want string
	}{
		// TODO: Add test cases.
		{name: "simple test",
			b:    DefaultDataBit{_tableName: "shreesha"},
			want: "shreesha",
		},
		{name: "simple empty bit name test",
			b:    DefaultDataBit{_tableName: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.BitType(); got != tt.want {
				t.Errorf("DefaultDataBit.BitName() = %v, want %v", got, tt.want)
			}
		})
	}
}
