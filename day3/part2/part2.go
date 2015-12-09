package main

import (
	"bufio"
	"fmt"
	"os"
)

type Santa struct {
	row, col int // All a Santa has to know is where he is
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	// with a santas array and a selector we can easily swap
	// between them while we iterate the instructions
	santa := &Santa{row: 0, col: 0}
	robot := &Santa{row: 0, col: 0}
	santas := []Santa{*santa, *robot}
	selector := 0

	grid := make(map[int]map[int]int)
	row, col := 0, 0

	for s.Scan() {
		line := s.Bytes()
		for i := 0; i < len(line); i++ {
			if i%2 == 0 {
				selector = 0
			} else {
				selector = 1
			}

			switch string(line[i]) {
			case "^":
				santas[selector].row++
			case "v":
				santas[selector].row--
			case ">":
				santas[selector].col++
			case "<":
				santas[selector].col--
			}

			row, col = santas[selector].row, santas[selector].col
			_, ok := grid[row]
			if !ok {
				grid[row] = make(map[int]int)
			}

			grid[row][col]++
		}
	}

	houses := 0
	for r, _ := range grid {
		for _, _ = range grid[r] {
			houses++
		}
	}

	fmt.Printf("%v houses get at least one present.\n", houses)
}
