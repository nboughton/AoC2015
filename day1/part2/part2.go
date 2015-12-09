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

	floor, character := 0, 0

	for s.Scan() {
		line := s.Bytes()
		for i := 0; i < len(line); i++ {
			if string(line[i]) == "(" {
				floor++
			}
			if string(line[i]) == ")" {
				floor--
			}

			if floor == -1 {
				fmt.Printf("Character position: %v\n", character)
				break
			}
			character++
		}
	}
}
