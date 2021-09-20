package main

import (
	"testing"
)

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		name string
		a, b int
		want int
	}{
		{"0_and_1", 0, 1, 0},
		{"1_and_0", 1, 0, 0},
		{"2_and_-2", 2, -2, -2},
		{"0_and_-1", 0, -1, -1},
		{"-1_and_0", -1, 0, -1},
	}

	for _, tt := range tests {
		// testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(tt.name, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
