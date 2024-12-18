// Look into using this library: https://stackoverflow.com/questions/23923383/evaluate-formula-in-go
package day07

import (
	"strconv"
	"strings"

	"github.com/thomasjazz/adventofcode/lib"
)

var (
	filename = "../../input/input-day07.txt"
)

type Equation struct {
	target uint64
	args   []uint64
}

func PartOne() uint64 {
	lines := lib.ReadInput(filename)
	equations, _ := ParseEquations(lines)
	var calibrationSum uint64 = 0

	for _, equation := range equations {
		result := evaluateEquation(equation, 0, 0)
		if result {
			calibrationSum += equation.target
		}
	}
	return calibrationSum
}

func PartTwo() uint64 {
	lines := lib.ReadInput(filename)
	equations, _ := ParseEquations(lines)
	var calibrationSum uint64 = 0

	for _, equation := range equations {
		result := evaluateEquation2(equation, 0, 0)
		if result {
			calibrationSum += equation.target
		}
	}
	return calibrationSum
}

func evaluateEquation(eq Equation, index int, currentValue uint64) bool {
	if index == len(eq.args) {
		return currentValue == eq.target
	}

	return evaluateEquation(eq, index+1, currentValue+eq.args[index]) || evaluateEquation(eq, index+1, currentValue*eq.args[index])
}

func evaluateEquation2(eq Equation, index int, currentValue uint64) bool {
	if index == len(eq.args) {
		return currentValue == eq.target
	}

	//concatenatedValue, _ := strconv.Atoi(strconv.Itoa(eq.args[index]) + strconv.Itoa(eq.args[index+1]))
	concatenatedValue := concatValue(currentValue, eq.args[index])

	return evaluateEquation2(eq, index+1, currentValue+eq.args[index]) ||
		evaluateEquation2(eq, index+1, currentValue*eq.args[index]) ||
		evaluateEquation2(eq, index+1, concatenatedValue)
}

func concatValue(a uint64, b uint64) uint64 {
	concatenatedValue, _ := strconv.ParseUint(strconv.Itoa(int(a))+strconv.Itoa(int(b)), 10, 64)
	return concatenatedValue
}

func ParseEquations(lines []string) ([]Equation, error) {
	equations := make([]Equation, 0)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		argStrings := strings.Split(parts[1], " ")[1:]
		target, _ := strconv.ParseUint(parts[0], 10, 64)

		args := make([]uint64, 0)
		for _, argStr := range argStrings {
			arg, _ := strconv.ParseUint(argStr, 10, 64)
			args = append(args, arg)
		}

		equation := Equation{
			target: target,
			args:   args,
		}
		equations = append(equations, equation)
	}
	return equations, nil
}
