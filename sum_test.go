package main

import "testing"

func TestSum(t *testing.T) {
	if Sum(5, 5) != 10 {
		t.Error("Expected value of sum 5 + 5 must equal 10")
	}
}
