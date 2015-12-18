package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var raceTime = 2503

type reindeer struct {
	name                                  string
	flySpeed, flyTime, restTime, distance int
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	longest := 0

	for s.Scan() {
		r := parseLine(s.Text())
		r.race(raceTime)
		if longest < r.distance {
			longest = r.distance
		}
		fmt.Printf("%v: %v\n", r.name, r.distance)
	}
	fmt.Printf("The winner has travelled %vkm\n", longest)
}

func (r *reindeer) race(t int) {
	c := 0
	for i := 0; i < t; i++ { // iterate each second
		if c < r.flyTime {
			// flying
			c++
			r.distance += r.flySpeed
		} else if c < r.flyTime+r.restTime {
			// resting
			if c == r.flyTime+r.restTime-1 {
				c = 0
			} else {
				c++
			}
		}
	}
}

func parseLine(s string) *reindeer {
	a := strings.Split(s, " ")
	n := a[0]

	fs, _ := strconv.Atoi(a[3])
	ft, _ := strconv.Atoi(a[6])
	rt, _ := strconv.Atoi(a[13])

	return &reindeer{name: n, flySpeed: fs, flyTime: ft, restTime: rt}
}
