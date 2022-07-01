package main_test

import (
	"testing"
)

func TestAddition(t *testing.T) {
	got := 2 + 3
	expected := 4
	if got != expected {
		t.Errorf("Did not get expected result.")
	}
}
