package main

import (
	"bufio"
	"fmt"
	"os"
)

type lightGrid [][]bool

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	// read grid for initial state
	grid := make(lightGrid, 100)
	i := 0
	for s.Scan() {
		for _, c := range s.Bytes() {
			if string(c) == "#" {
				grid[i] = append(grid[i], true)
			} else {
				grid[i] = append(grid[i], false)
			}
		}
		i++
	}

	printGrid(grid)
	for i := 0; i < 100; i++ {
		grid = step(grid)
		fmt.Printf("Step %v\n", i)
		printGrid(grid)
	}

	lightsOn := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] {
				lightsOn++
			}
		}
	}
	fmt.Printf("%v lights are left on\n", lightsOn)
}

func step(g lightGrid) lightGrid {
	var ng = make(lightGrid, len(g))
	for i := 0; i < len(g); i++ {
		ng[i] = make([]bool, len(g))
	}

	for row := range g {
		for col := range g[row] {
			ng[row][col] = testLight(row, col, g)
		}
	}

	return ng
}

func testLight(r, c int, g lightGrid) bool {
	on := g[r][c]
	surOn := 0

	for row := r - 1; row <= r+1; row++ {
		for col := c - 1; col <= c+1; col++ {
			if row >= 0 && col >= 0 && row < len(g) && col < len(g) {
				if row == r && col == c {
					continue
				} else if g[row][col] {
					surOn++
				}
			}
		}
	}

	if on {
		if surOn == 2 || surOn == 3 {
			return true
		} else {
			return false
		}
	}

	if !on && surOn == 3 {
		return true
	}
	return false
}

func printGrid(g lightGrid) {
	for row := range g {
		for col := range g[row] {
			if g[row][col] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
