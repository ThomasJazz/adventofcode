package day1

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := 1150
	// Add tests here
	timesHitZero := PartOne()
	println(timesHitZero)

	if expected != timesHitZero {
		t.Fatal("Expected output was ", expected, "but actual was ", timesHitZero)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 6738
	timesHitZero := PartTwo()
	println(timesHitZero)

	if expected != timesHitZero {
		t.Fatal("Expected output was ", expected, "but actual was ", timesHitZero)
	}
}
