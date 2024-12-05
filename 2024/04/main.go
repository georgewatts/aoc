package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

const (
	X        = 88
	M        = 77
	A        = 65
	S        = 83
	NEW_LINE = 10
)

func checkCurrentLine(pos int, chars []byte) int {
	// check fwd
	retVal := 0
	fwd := chars[pos+1 : pos+4]
	if string(fwd) == "MAS" {
		retVal++
	}

	// check back only if we can
	if (pos - 3) > 0 {
		back := chars[pos-3 : pos]
		if string(back) == "SAM" {
			retVal++
		}
	}

	return retVal
}

func checkBelow(pos int, lines [][]byte) int {
	retVal := 0

	belowChars := []byte{}
	for i := 0; i < 3; i++ {
		belowChars = append(belowChars, lines[i][pos])
	}
	if string(belowChars) == "MAS" {
		retVal++
	}

	fwdChars := []byte{}
	for i := 0; i < 3; i++ {
		if (pos + i) >= len(lines[i]) {
			break
		}

		fwdChars = append(fwdChars, lines[i][pos+i])
	}
	if string(fwdChars) == "MAS" {
		retVal++
	}

	// check back diag
	backChars := []byte{}
	for i := 1; i < 4; i++ {
		if (pos - i) <= 0 {
			break
		}

		backChars = append(backChars, lines[i-1][pos-i])
	}

	if string(backChars) == "MAS" {
		retVal++
	}

	return retVal
}

func checkAbove(pos int, lines [][]byte) int {
	retVal := 0

	aboveChars := []byte{}
	for i := 0; i < 3; i++ {
		aboveChars = append(aboveChars, lines[i][pos])
	}
	if string(aboveChars) == "MAS" {
		fmt.Println("ABOVE POS:", pos)
		retVal++
	}

	fwdChars := []byte{}
	for i := 0; i < 3; i++ {
		if (pos + i + 1) >= len(lines[i]) {
			break
		}

		fwdChars = append(fwdChars, lines[i][pos+i+1])
	}
	if string(fwdChars) == "MAS" {
		fmt.Println("FWD DIAG POS:", pos)
		retVal++
	}

	// check back diag
	backChars := []byte{}
	for i := 0; i < 3; i++ {
		if (pos - i - 1) < 0 {
			break
		}

		backChars = append(backChars, lines[i][pos-i-1])
	}

	if string(backChars) == "MAS" {
		fmt.Println("BACK DIAG POS:", pos)
		retVal++
	}

	return retVal
}

func main() {
	grid := [][]byte{}

	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	count := 0
	maxBelow := len(grid) - 3

	for line, v := range grid {
		fmt.Println("LINE:", line)
		for pos, v2 := range v {
			if v2 == X {
				count += checkCurrentLine(pos, v)
				if (line + 3) < maxBelow {
					count += checkBelow(pos, grid[line+1:line+4])
					// checkBelow(pos, v[line+1:line+4])
					// fmt.Println("Line", line, "Pos", pos)
				}
				if (line) > 2 {
					lines := slices.Clone(grid[line-3 : line])
					slices.Reverse(lines)
					count += checkAbove(pos, lines)
				}
			}
		}
	}

	fmt.Printf("count: %v\n", count)
}
