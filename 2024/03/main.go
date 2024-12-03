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

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	data, _ := io.ReadAll(inputFile)
	str := string(data)

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	occurences := re.FindAllString(str, -1)

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

	fmt.Printf("total: %v\n", total)
}
