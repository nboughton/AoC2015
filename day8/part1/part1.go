package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var hex = regexp.MustCompile(`\\x[a-z0-9]{2}`)
var dbs = regexp.MustCompile(`\\[\\"]{1}`)

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	totalChars := 0
	totalMem := 0
	for s.Scan() {
		line := s.Text()
		totalChars += len([]byte(line))
		totalMem += getMem(line)
		fmt.Printf("c: %v, m: %v\n", len([]byte(line)), getMem(line))
	}

	fmt.Printf("%v\n", totalChars-totalMem)
}

func getMem(s string) int {
	mem := len([]byte(s)) - 2
	mem -= len(hex.FindAllString(s, -1)) * 3
	mem -= len(dbs.FindAllString(s, -1))

	return mem
}
