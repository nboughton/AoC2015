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

	result := 0
	for s.Scan() {
		line := s.Text()
		result += len([]byte(line)) - getMem(line)
	}

	fmt.Printf("%v\n", result)
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
