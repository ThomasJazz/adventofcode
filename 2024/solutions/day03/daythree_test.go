package day03

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := 189527826
	response := PartOne()
	println(response)

	if response != expected {
		t.Fatalf(`Expected value: %d, Actual value: %d`, expected, response)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 604
	response := PartTwo()
	println(response)

	if response != expected {
		t.Fatalf(`Expected value: %d, Actual value: %d`, expected, response)
	}
}
