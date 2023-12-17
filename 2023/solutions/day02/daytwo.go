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

func CubeConundrum() {
	lines := lib.ReadInput(filepath)
	gameIndexSum := 0

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
	}

	println(gameIndexSum)
}
