package main

import "testing"

func TestAdd(t *testing.T) {

	if Add(4, 5) != 9 {
		t.Error("error")
	}

}
func BenchmarkAdd(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Add(4, 5)
	}
}
