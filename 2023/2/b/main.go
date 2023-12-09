package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Set struct {
	red   int
	blue  int
	green int
}

type Game struct {
	id   int
	sets []Set
}

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func parseSet(rawSet string) Set {
	red, blue, green := 0, 0, 0
	scores := strings.Split(rawSet, ", ")

	for _, v := range scores {
		score := strings.Split(v, " ")
		switch score[1] {
		case "red":
			red, _ = strconv.Atoi(score[0])
		case "blue":
			blue, _ = strconv.Atoi(score[0])
		case "green":
			green, _ = strconv.Atoi(score[0])
		}
	}

	return Set{
		red, blue, green,
	}
}

func parseGame(rawId string, rawGameString string) Game {
	sets := []Set{}
	parsedID, _ := strconv.Atoi(rawId)
	rawSets := strings.Split(rawGameString, "; ")

	for _, rawSet := range rawSets {
		sets = append(sets, parseSet(rawSet))
	}

	return Game{
		id:   parsedID,
		sets: sets,
	}
}

func (g Game) minSetValues() Set {
	red, green, blue := 0, 0, 0
	for _, s := range g.sets {
		if red < s.red {
			red = s.red
		}
		if green < s.green {
			green = s.green
		}
		if blue < s.blue {
			blue = s.blue
		}
	}

	return Set{
		red, green, blue,
	}
}

func (s Set) power() int {
	return s.red * s.green * s.blue
}

func main() {
	inputFile, err := os.Open("input.txt")

	if err != nil {
		panic("open error")
	}

	scanner := bufio.NewScanner(inputFile)

	scanner.Scan()
	line := scanner.Text()

	games := []Game{}

	for len(line) != 0 {
		regex := regexp.MustCompile(`Game\s(?P<index>\d+):\s(?P<games>.*)`)
		rawGame := regex.FindStringSubmatch(line)
		// fmt.Printf("rawGame: %v\n", rawGame)
		// for i, v := range rawGame {
		// 	fmt.Printf("rawGame[%d]: %v\n", i, v)
		// }

		games = append(games, parseGame(rawGame[1], rawGame[2]))
		scanner.Scan()
		line = scanner.Text()
	}

	sum := 0

	for _, g := range games {
		sum += g.minSetValues().power()
	}

	fmt.Printf("sum: %v\n", sum)
}
