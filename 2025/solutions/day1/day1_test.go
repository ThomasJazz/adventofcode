package day1

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := 1102
	// Add tests here
	timesHitZero := PartOne()
	println(timesHitZero)

	if expected != timesHitZero {
		t.Fatal("Expected output was ", expected, "but actual was ", timesHitZero)
	}
}

func TestPartTwo(t *testing.T) {
	timesHitZero := PartTwo()
	println(timesHitZero)
}
