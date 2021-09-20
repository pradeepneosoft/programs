package main

import (
	"testing"
)

func Testsearchinslice(t *testing.T) {
	if searchinslice(3, 4, 5, 76, 7, 3, 2, 4, 5, 6, 75, 4, 3, 21, 2, 3, 56, 5, 78) != 6 {
		t.Error("failed")
	}

}
func Testsearchinslice1(t *testing.T) {
	if searchinslice1(3, 4, 5, 76, 7, 3, 2, 4, 5, 6, 75, 4, 3, 21, 2, 3, 56, 5, 78) != 2 {
		t.Error("failed")
	}

}
