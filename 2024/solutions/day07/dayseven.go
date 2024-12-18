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
	target int
	args   []int
}

func PartOne() int {
	lines := lib.ReadInput(filename)
	equations, _ := ParseEquations(lines)
	calibrationSum := 0

	for _, equation := range equations {
		result := evaluateEquation(equation, 0, 0)
		if result {
			calibrationSum += equation.target
		}
	}
	return calibrationSum
}

func evaluateEquation(eq Equation, index int, currentValue int) bool {
	if index == len(eq.args) {
		return currentValue == eq.target
	}

	return evaluateEquation(eq, index+1, currentValue+eq.args[index]) || evaluateEquation(eq, index+1, currentValue*eq.args[index])
}

func ParseEquations(lines []string) ([]Equation, error) {
	equations := make([]Equation, 0)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		argStrings := strings.Split(parts[1], " ")[1:]
		target, _ := strconv.Atoi(parts[0])

		args := make([]int, 0)
		for _, argStr := range argStrings {
			arg, _ := strconv.Atoi(argStr)
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
