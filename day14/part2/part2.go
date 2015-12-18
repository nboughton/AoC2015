package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var raceTime = 2503

type reindeer struct {
	name                                                  string
	flySpeed, flyTime, restTime, distance, points, sCount int
}

type leader struct {
	name   string
	points int
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	racers := make(map[string]*reindeer)

	for s.Scan() {
		r := parseLine(s.Text())
		racers[r.name] = r
	}
	lead := race(racers, raceTime)
	fmt.Printf("The winner(%v) has %v points\n", lead.name, lead.points)
}

func race(r map[string]*reindeer, t int) *leader {
	l := &leader{}
	for i := 0; i < t; i++ { // iterate each second
		for rd := range r {
			if r[rd].sCount < r[rd].flyTime {
				// flying
				r[rd].sCount++
				r[rd].distance += r[rd].flySpeed
			} else if r[rd].sCount < r[rd].flyTime+r[rd].restTime {
				// resting
				if r[rd].sCount == r[rd].flyTime+r[rd].restTime-1 {
					r[rd].sCount = 0
				} else {
					r[rd].sCount++
				}
			}
		}

		// Let's see who's in the lead and award points
		d := []int{}
		for rd := range r {
			d = append(d, r[rd].distance)
		}
		sort.Ints(d)

		for rd := range r {
			if r[rd].distance == d[len(d)-1] {
				r[rd].points++
				if r[rd].points > l.points {
					l.name = rd
					l.points = r[rd].points
				}
			}
		}
	}
	return l
}

func parseLine(s string) *reindeer {
	a := strings.Split(s, " ")
	n := a[0]

	fs, _ := strconv.Atoi(a[3])
	ft, _ := strconv.Atoi(a[6])
	rt, _ := strconv.Atoi(a[13])

	return &reindeer{name: n, flySpeed: fs, flyTime: ft, restTime: rt}
}
