package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var cal = make(map[string][]string)
var input string

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		line := s.Text()
		if strings.Contains(line, "=>") {
			p := strings.Split(line, " => ")
			cal[p[0]] = append(cal[p[0]], p[1])
		} else if len(line) > 1 {
			input = line
		}
	}

	// I'm a terrible person.
	tree("e", 0)
}

func tree(st string, step int) {
	for mol, _ := range cal {
		t := regexp.MustCompile(`(` + mol + `)`)
		n := t.FindAllIndex([]byte(st), -1)

		for i := 0; i < len(n); i++ {
			for _, sub := range cal[mol] {
				st = applySub([]byte(st), []byte(sub), n[i])
				if len(st) <= len(input) && st != input {
					fmt.Printf("step %v, mol: %v\n", step, st)
					step++
					tree(st, step)
				} else if st == input {
					fmt.Printf("%v steps to %v\n", step, input)
					os.Exit(0)
				}
			}
		}
	}
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
