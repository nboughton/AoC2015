package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	cal := make(map[string][]string)
	res := make(map[string]int)

	input := ""
	for s.Scan() {
		line := s.Text()
		if strings.Contains(line, "=>") {
			p := strings.Split(line, " => ")
			cal[p[0]] = append(cal[p[0]], p[1])
		} else if len(line) > 1 {
			input = line
		}
	}

	for mol, _ := range cal {
		t := regexp.MustCompile(`(` + mol + `)`)
		n := t.FindAllIndex([]byte(input), -1)

		for i := 0; i < len(n); i++ {
			for _, sub := range cal[mol] {
				res[applySub([]byte(input), []byte(sub), n[i])]++
			}
		}
	}

	fmt.Printf("%v distinct molecules created\n", len(res))
}

func applySub(in, sub []byte, i []int) string {
	out := []byte{}

	for index, value := range in {
		if index == i[0] {
			out = append(out, sub...)
		} else if index > i[0] && index < i[1] {
			continue
		} else {
			out = append(out, value)
		}
	}

	return string(out)
}
