package ipTools

import (
	"testing"
)

func TestIP_IsIPv4(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"0.0.0.0/0", false},
		{"192.168.1.1/20", true},
	}

	for _, test := range tests {
		addr := &IP{
			address: test.in,
		}
		got, _ := addr.IsIPv4()
		if test.out != got {
			t.Errorf("in: %v, want: %v, but got: %v", test.in, test.out, got)
		}
	}
}
