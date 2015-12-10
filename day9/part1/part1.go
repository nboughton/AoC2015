package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	distances := make(map[string]map[string]int)

	for s.Scan() {
		from, to, distance := ParseLine(s.Text())

		_, ok := distances[from]
		if !ok {
			distances[from] = make(map[string]int)
		}

		distances[from][to] = distance
	}

	/*
		distance := 1000 // Anything will be shorter than this
		for from, _ := range distances {
			for to, _ := range distances[from] {
				for dest, _ := range distances[to] {
					d := distances[from][to] + distances[to][dest]
					fmt.Printf("%v -> %v -> %v: %v\n", from, to, dest, d)
					if d < distance {
						distance = d
					}
				}
			}
		}
	*/
	fmt.Printf("The shortest route is %v\n")
}

func ParseLine(s string) (string, string, int) {
	atoms := strings.Split(s, " ")
	from, to := atoms[0], atoms[2]
	distance, _ := strconv.ParseInt(atoms[4], 10, 32)

	return from, to, int(distance)
}
