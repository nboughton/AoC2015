package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var ptrn = regexp.MustCompile(`\\(x[a-z0-9]{2}|[\\"]{1})`)

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	result1 := 0
	result2 := 0
	for s.Scan() {
		line := s.Text()
		result1 += len([]byte(line)) - getMem(line)
		result2 += len([]byte(fmt.Sprintf("%q", line))) - len([]byte(line))
	}

	fmt.Printf("Part 1: %v\nPart 2: %v\n", result1, result2)
}

func getMem(s string) int {
	trim := string([]byte(s)[1 : len(s)-1])
	mem := len(trim)

	m := ptrn.FindAllString(trim, -1)
	for i := 0; i < len(m); i++ {
		if strings.Contains(m[i], "x") {
			mem -= 3
		} else {
			mem--
		}
	}

	return mem
}
