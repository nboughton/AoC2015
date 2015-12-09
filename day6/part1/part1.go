package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	// Create our light grid
	gridSize := 1000
	grid := make([][]bool, gridSize)
	for i := 0; i < gridSize; i++ {
		grid[i] = make([]bool, gridSize)
	}

	for s.Scan() {
		line := s.Text()
		action, startRow, startCol, endRow, endCol := ParseLine(line)

		for row := startRow; row <= endRow; row++ {
			for col := startCol; col <= endCol; col++ {
				switch action {
				case "on":
					grid[row][col] = true
				case "off":
					grid[row][col] = false
				case "toggle":
					if grid[row][col] {
						grid[row][col] = false
					} else {
						grid[row][col] = true
					}
				}
			}
		}
	}

	lightsOn := 0
	for row, _ := range grid {
		for col, _ := range grid[row] {
			if grid[row][col] == true {
				lightsOn++
			}
		}
	}

	fmt.Printf("%v lights are turned on\n", lightsOn)
}

func ParseLine(line string) (string, int, int, int, int) {
	action, start, end := "", "", ""
	startRow, startCol, endRow, endCol := 0, 0, 0, 0
	atoms := strings.Split(line, " ")

	if strings.Contains(line, "toggle") {
		action = "toggle"
		start = atoms[1]
		end = atoms[3]
	} else {
		action = atoms[1]
		start = atoms[2]
		end = atoms[4]
	}

	atoms = strings.Split(start, ",")
	sr, _ := strconv.ParseInt(atoms[0], 10, 32)
	sc, _ := strconv.ParseInt(atoms[1], 10, 32)

	atoms = strings.Split(end, ",")
	er, _ := strconv.ParseInt(atoms[0], 10, 32)
	ec, _ := strconv.ParseInt(atoms[1], 10, 32)

	startRow, startCol, endRow, endCol = int(sr), int(sc), int(er), int(ec)
	return action, startRow, startCol, endRow, endCol
}
