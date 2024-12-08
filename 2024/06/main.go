package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	LEFT  = "<"
	RIGHT = ">"
	UP    = "^"
	DOWN  = "v"
)

type Coord struct {
	x int
	y int
}

type Layout [][]string

type RoomMap struct {
	GuardOrientation string
	GuardPos         Coord
	Locs             map[Coord]bool
	Layout
}

func (l Layout) String() string {
	retVal := "\n"
	for _, v := range l {
		retVal += "["
		retVal += strings.Join(v, " ")
		retVal += "]"
		retVal += "\n"
	}
	return retVal
}

func (m RoomMap) GetNextCoord() string {
	switch m.GuardOrientation {
	case LEFT:
		if m.GuardPos.x-1 < 0 {
			fmt.Println("Can't move LEFT")
			return "-1"
		}
		return m.Layout[m.GuardPos.y][m.GuardPos.x-1]
	case RIGHT:
		if m.GuardPos.x+1 > len(m.Layout[0]) {
			fmt.Println("Can't move RIGHT")
			return "-1"
		}
		return m.Layout[m.GuardPos.y][m.GuardPos.x+1]
	case UP:
		if m.GuardPos.y-1 < 0 {
			fmt.Println("Can't move UP")
			return "-1"
		}
		return m.Layout[m.GuardPos.y-1][m.GuardPos.x]
	case DOWN:
		if m.GuardPos.y+1 >= len(m.Layout) {
			fmt.Println("Can't move DOWN")
			return "-1"
		}
		return m.Layout[m.GuardPos.y+1][m.GuardPos.x]
	}

	fmt.Println("OOPPSSIIIEEE")
	return "-1"
}

func (m *RoomMap) MoveGuard() {
	nextCoord := m.GetNextCoord()
	fmt.Printf("m.GuardPos: %v\n", m.GuardPos)

	if nextCoord == "-1" {
		fmt.Println("Early return")
		return
	}

	if nextCoord == "#" {
		switch m.GuardOrientation {
		case LEFT:
			m.GuardOrientation = UP
		case UP:
			m.GuardOrientation = RIGHT
		case RIGHT:
			m.GuardOrientation = DOWN
		case DOWN:
			m.GuardOrientation = LEFT
		}
	}

	if nextCoord != "#" {
		switch m.GuardOrientation {
		case LEFT:
			m.GuardPos = Coord{x: m.GuardPos.x - 1, y: m.GuardPos.y}
		case RIGHT:
			m.GuardPos = Coord{x: m.GuardPos.x + 1, y: m.GuardPos.y}
		case UP:
			m.GuardPos = Coord{x: m.GuardPos.x, y: m.GuardPos.y - 1}
		case DOWN:
			m.GuardPos = Coord{x: m.GuardPos.x, y: m.GuardPos.y + 1}
		}
		m.Locs[m.GuardPos] = true
	}

	m.MoveGuard()
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputFile)
	roomMap := RoomMap{}
	roomMap.Locs = make(map[Coord]bool)
	roomMap.GuardOrientation = "^"
	lineNum := 0

	for scanner.Scan() {
		line := scanner.Text()

		chars := strings.Split(line, "")

		if strings.Contains(line, "^") {
			xPos := strings.IndexAny(line, "^")
			coord := Coord{x: xPos, y: lineNum}
			roomMap.GuardPos = coord
			roomMap.Locs[coord] = true
		}

		roomMap.Layout = append(roomMap.Layout, chars)
		lineNum++
	}

	// roomMap.GuardPos.y = len(roomMap.Layout) - 1 - roomMap.GuardPos.y
	fmt.Printf("roomMap.Layout: %v\n", roomMap.Layout)
	fmt.Printf("start GuardPos: %v\n", roomMap.GuardPos)

	roomMap.MoveGuard()

	fmt.Printf("roomMap.GuardPos: %v\n", roomMap.GuardPos)
	fmt.Printf("len(roomMap.Locs): %v\n", len(roomMap.Locs))
}
