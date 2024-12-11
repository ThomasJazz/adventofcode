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
	expected int
	args     []int
}

func PartOne() int {
	lines := lib.ReadInput(filename)
	equations, _ := ParseEquations(lines)

	for _, equation := range equations {
		println(CanMakeExpected(equation))
	}
	return 0
}

func CanMakeExpected(equation Equation) bool {
	expressions := GenerateAllExpressions(equation.args)

	println(expressions)
	return false
}

func GenerateAllExpressions(args []int) []string {
	//operands := []string{"+", "*"}
	expressions := make([]string, 0)

	for i := 0; i < len(args); i++ {
		expression := buildCombosRecursive(args, i, strconv.Itoa(args[i]), "+")
		expressions = append(expressions, expression)
		expression = buildCombosRecursive(args, i, strconv.Itoa(args[i]), "*")
		expressions = append(expressions, expression)
	}

	return expressions
}

func buildCombosRecursive(args []int, index int, expression string, operand string) string {
	if index == len(args) {
		return expression
	}

	if operand == "*" {
		return buildCombosRecursive(args, index+1, expression+operand+strconv.Itoa(args[index]), "+")
	} else {
		return buildCombosRecursive(args, index+1, expression+operand+strconv.Itoa(args[index]), "*")
	}
}

func ParseEquations(lines []string) ([]Equation, error) {
	equations := make([]Equation, 0)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		argStrings := strings.Split(parts[1], " ")[1:]
		expected, _ := strconv.Atoi(parts[0])

		args := make([]int, 0)
		for _, argStr := range argStrings {
			arg, _ := strconv.Atoi(argStr)
			args = append(args, arg)
		}

		equation := Equation{
			expected: expected,
			args:     args,
		}
		equations = append(equations, equation)
	}
	return equations, nil
}
