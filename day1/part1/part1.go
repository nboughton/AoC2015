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

	var floor int
	var instructions []byte

	for s.Scan() {
		instructions = append(instructions, s.Bytes()...)
	}

	for i := 0; i < len(instructions); i++ {
		if string(instructions[i]) == "(" {
			floor++
		}
		if string(instructions[i]) == ")" {
			floor--
		}
	}

	fmt.Printf("Floor: %v\n", floor)
}
