package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	grid := make(map[int]map[int]int)
	row, col := 0, 0

	for s.Scan() {
		line := s.Bytes()
		for i := 0; i < len(line); i++ {
			switch string(line[i]) {
			case "^":
				row++
			case "v":
				row--
			case ">":
				col++
			case "<":
				col--
			}

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
