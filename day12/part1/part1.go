package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var isInt = regexp.MustCompile(`([-0-9]+)`)

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	sum := 0
	for s.Scan() {
		line := s.Text()
		n := isInt.FindAllString(line, -1)
		for i := 0; i < len(n); i++ {
			num, _ := strconv.Atoi(n[i])
			sum += num
		}
	}
	fmt.Printf("%v\n", sum)
}
