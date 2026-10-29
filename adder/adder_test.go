package main

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(1, 2)
	expected := 3

	if sum != expected {
		t.Errorf("expected: %d, got: %d", expected, sum)
	}
}
