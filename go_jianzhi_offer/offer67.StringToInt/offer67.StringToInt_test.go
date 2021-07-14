package offer

import (
	"fmt"
	"testing"
)

func Test_strToInt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{"42"}, 42},
		{"1", args{" -42"}, -42},
		{"2", args{"4193 with words"}, 4193},
		{"3", args{"words and 987"}, 0},
		{"4", args{"-91283472332"}, -2147483648},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strToInt(tt.args.str); got != tt.want {
				t.Errorf("strToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PrintInt(t *testing.T) {
	fmt.Printf("INT_MAX=%d\nINT_MIN=%d\n", 1<<31-1, -1<<31)
}
