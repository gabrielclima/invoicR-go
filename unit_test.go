package main

import (
	"testing"
)

func Test(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.4 {
		t.Error("Expected 1.5, got ", v)
	}
}
