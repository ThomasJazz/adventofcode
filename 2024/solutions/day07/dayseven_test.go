package day07

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := 0
	actual := PartOne()
	if actual != expected {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}
