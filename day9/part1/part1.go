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
	places := []string{}
	for s.Scan() {
		from, to, distance := parseLine(s.Text())

		_, ok := distances[from]
		if !ok {
			distances[from] = make(map[string]int)
		}

		distances[from][to] = distance
	}

	fmt.Printf("places %v\n")
}

func addPlace(p string, places *[]string) {
	for i := 0; i < len(*places); i++ {

	}
}

func tryPath(start, end string, d map[string]map[string]int) int {
	//path := start

	return 0
}

func parseLine(s string) (string, string, int) {
	atoms := strings.Split(s, " ")
	from, to := atoms[0], atoms[2]
	distance, _ := strconv.ParseInt(atoms[4], 10, 32)

	return from, to, int(distance)
}
