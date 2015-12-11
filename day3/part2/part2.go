package main

import (
	"bufio"
	"fmt"
	"os"
)

type santa struct {
	row, col int // All a Santa has to know is where he is
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	// with a santas array and a sel we can easily swap
	// between them while we iterate the instructions
	sta := &santa{row: 0, col: 0}
	rbt := &santa{row: 0, col: 0}
	santas := []santa{*sta, *rbt}
	sel := 0

	grid := make(map[int]map[int]int)
	row, col := 0, 0

	for s.Scan() {
		line := s.Bytes()
		for i := 0; i < len(line); i++ {
			if i%2 == 0 {
				sel = 0
			} else {
				sel = 1
			}

			switch string(line[i]) {
			case "^":
				santas[sel].row++
			case "v":
				santas[sel].row--
			case ">":
				santas[sel].col++
			case "<":
				santas[sel].col--
			}

			row, col = santas[sel].row, santas[sel].col
			_, ok := grid[row]
			if !ok {
				grid[row] = make(map[int]int)
			}

			grid[row][col]++
		}
	}

	houses := 0
	for r := range grid {
		for _ = range grid[r] {
			houses++
		}
	}

	fmt.Printf("%v houses get at least one present.\n", houses)
}
