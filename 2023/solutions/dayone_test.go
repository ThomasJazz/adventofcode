package solutions

import (
	"testing"
)

func TestCalibrationNumbersTwo(t *testing.T) {
	expected := 18
	input := "one23testoneight"
	calibrationNum, _ := FindCalibrationNumbersTwo(input)

	if calibrationNum != expected {
		t.Fatalf(`Expected value: %d, Actual value: %d`, expected, calibrationNum)
	}
}
