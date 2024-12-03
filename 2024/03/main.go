package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Mul struct {
	a int
	b int
}

func processMultipliers(occurences []string) int {
	multipliers := []Mul{}

	for _, v := range occurences {
		formatted := strings.TrimLeft(v, "mul(")
		formatted = strings.TrimRight(formatted, ")")
		ints := strings.Split(formatted, ",")
		a, _ := strconv.Atoi(ints[0])
		b, _ := strconv.Atoi(ints[1])
		multipliers = append(multipliers, Mul{a, b})
	}

	total := 0
	for _, m := range multipliers {
		total += m.a * m.b
	}

	return total
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	data, _ := io.ReadAll(inputFile)
	str := string(data)

	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don\'t\(\)`)
	occurences := re.FindAllString(str, -1)

	filteredOccurences := []string{}
	ignoringOccurence := false
	for _, v := range occurences {
		if v == "do()" {
			ignoringOccurence = false
		} else if v == "don't()" {
			ignoringOccurence = true
		} else {
			if !ignoringOccurence {
				filteredOccurences = append(filteredOccurences, v)
			}
		}
	}

	total := processMultipliers(filteredOccurences)
	fmt.Printf("total: %v\n", total)
}
