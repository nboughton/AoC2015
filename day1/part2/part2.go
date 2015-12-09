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

	var floor, character int
	var instructions []byte

	for s.Scan() {
		instructions = append(instructions, s.Bytes()...)
	}

	character = 1
	for i := 0; i < len(instructions); i++ {
		if string(instructions[i]) == "(" {
			floor++
		}
		if string(instructions[i]) == ")" {
			floor--
		}

		if floor == -1 {
			fmt.Printf("Chararcter position: %v\n", character)
			break
		}
		character++
	}
}
