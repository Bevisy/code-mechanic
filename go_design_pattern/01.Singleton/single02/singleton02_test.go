package main

import (
	"reflect"
	"testing"
)

func Test_getInstance(t *testing.T) {
	tests := []struct {
		name string
		want *single
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
