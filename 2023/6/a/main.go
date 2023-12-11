package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Race struct {
	time          int
	recordDist    int
	winningAccels int
}

func (r Race) calculateDistanceForAccelTime(accelTime int) int {
	return accelTime * (r.time - accelTime)
}

func getRaces(raceValues []string) []Race {
	times := raceValues[0 : len(raceValues)/2]
	distances := raceValues[len(raceValues)/2:]

	races := []Race{}

	// fmt.Printf("times: %v\n", times)
	// fmt.Printf("distances: %v\n", distances)

	for i, v := range times {
		time, _ := strconv.Atoi(v)
		dist, _ := strconv.Atoi(distances[i])
		races = append(races, Race{time: time, recordDist: dist, winningAccels: 0})
	}

	processWinningAccels(races)

	return races
}

func processWinningAccels(races []Race) {
	for i := 0; i < len(races); i++ {
		min, max := 0, 0
		for i2 := 0; i2 < races[i].recordDist; i2++ {
			if races[i].calculateDistanceForAccelTime(i2) > races[i].recordDist {
				min = i2
				break
			}
		}

		for i3 := races[i].recordDist; i3 > min; i3-- {
			if races[i].calculateDistanceForAccelTime(i3) > races[i].recordDist {
				max = i3
				break
			}
		}

		races[i].winningAccels = max - min + 1
	}
}

func main() {
	inputFile, err := os.Open("input.txt")

	if err != nil {
		panic("open error")
	}

	scanner := bufio.NewScanner(inputFile)

	scanner.Scan()
	line := scanner.Text()

	raceValues := []string{}

	for len(line) != 0 {
		raceValues = append(raceValues, regexp.MustCompile(`\d+`).FindAllString(line, -1)...)

		scanner.Scan()
		line = scanner.Text()
	}

	races := getRaces(raceValues)

	fmt.Printf("races: %v\n", races)

	result := 1

	for _, r := range races {
		result *= r.winningAccels
	}

	fmt.Printf("result: %v\n", result)
}
