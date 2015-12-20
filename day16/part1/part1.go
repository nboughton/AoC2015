package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var getNum = regexp.MustCompile(`([0-9]+)`)
var getVal = regexp.MustCompile(`([a-z]+: [0-9]+)`)
var sue = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		aunt := parseLine(s.Text())
		match := true
		for k := range aunt {
			if k != "num" && sue[k] != aunt[k] {
				match = false
				break
			}
		}
		if match {
			fmt.Printf("%v\n", aunt["num"])
		}
	}
}

func parseLine(s string) map[string]int {
	a := make(map[string]int)

	n := getNum.FindAllString(s, 1)
	a["num"], _ = strconv.Atoi(n[0])

	n = getVal.FindAllString(s, -1)

	p := []string{}
	v := 0
	for i := 0; i < len(n); i++ {
		p = strings.Split(n[i], ": ")
		v, _ = strconv.Atoi(p[1])
		a[p[0]] = v
	}

	return a
}
