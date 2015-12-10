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

	result := 0
	for s.Scan() {
		line := s.Text()
		result += len([]byte(line)) - getMem(line)
	}

	fmt.Printf("%v\n", result)
}

func getMem(s string) int {
	mem := len([]byte(s)) - 2
	mem -= len(hex.FindAllString(s, -1)) * 3
	mem -= len(dbs.FindAllString(s, -1))

	return mem
}
