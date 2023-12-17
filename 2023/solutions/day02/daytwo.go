// https://adventofcode.com/2023/day/2
package day02

import (
	"strconv"
	"strings"

	"github.com/thomasjazz/adventofcode/lib"
)

var (
	filepath   string         = "../../input/day_2_input.txt"
	cubesInBag map[string]int = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

func ParseGame(game string) []map[string]int {
	rounds := strings.Split(game, ";")
	var gameMap []map[string]int

	for i, round := range rounds {
		colors := strings.Split(round, ",")
		gameMap = append(gameMap, map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}) // Initialize game map

		for _, color := range colors {
			info := strings.Split(color, " ")
			gameMap[i][info[2]], _ = strconv.Atoi(info[1])
		}
	}

	return gameMap
}

func GetMinimumColorPower(game []map[string]int) int {
	var (
		minRed   = 0
		minBlue  = 0
		minGreen = 0
	)

	for _, round := range game {
		minRed = max(minRed, round["red"])
		minBlue = max(minBlue, round["blue"])
		minGreen = max(minGreen, round["green"])
	}

	return minBlue * minRed * minGreen
}

func CubeConundrum() {
	lines := lib.ReadInput(filepath)
	gameIndexSum := 0
	minPowerSum := 0

	for i, line := range lines {
		gameStr := strings.Split(line, ":")[1]
		game := ParseGame(gameStr)
		validGame := true

		for _, round := range game {
			if round["red"] > cubesInBag["red"] || round["blue"] > cubesInBag["blue"] || round["green"] > cubesInBag["green"] {
				validGame = false
			}
		}

		if validGame {
			gameIndexSum += (i + 1)
		}

		minPowerSum += GetMinimumColorPower(game)
	}

	println("Game index sum: %d", gameIndexSum)
	println("Minimum power sum: %d", minPowerSum)
}
