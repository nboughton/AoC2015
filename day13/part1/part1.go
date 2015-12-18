package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type hIndex struct {
	name string
	rel  map[string]int
}

// hIndexes is a global reference for seats and happinesss between seats
var hIndexes = make(map[string]*hIndex)

type seat struct {
	name, path string
	nextTo     map[string]*seat
}

type seatingPlan struct {
	happiness int
	path      string
}

var happiest = &seatingPlan{happiness: 0}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		n, happiness := parseLine(s.Text())

		_, ok := hIndexes[n[0]]
		if !ok {
			hIndexes[n[0]] = new(hIndex)
			hIndexes[n[0]].rel = make(map[string]int)
		}
		hIndexes[n[0]].rel[n[1]] = happiness
	}

	for a := range hIndexes {
		n := newSeat(a, a)
		n.addSeats()
	}

	fmt.Printf("happiest route: %v, happiness: %v\n", happiest.path, happiest.happiness)
}

func newSeat(name, path string) *seat {
	n := new(seat)
	n.nextTo = make(map[string]*seat)
	n.name = name
	n.incrementPath(path)

	return n
}

func (n *seat) incrementPath(path string) {
	n.path += " " + path + " "
	n.path = strings.TrimSpace(n.path)
}

func (n *seat) addSeats() {
	seatsAdded := 0
	for p := range hIndexes {
		if !strings.Contains(n.path, p) {
			_, ok := n.nextTo[p]
			if !ok {
				n.nextTo[p] = newSeat(p, n.path)
			}

			n.nextTo[p].incrementPath(p)
			n.nextTo[p].addSeats()
			seatsAdded++
		}
	}

	if seatsAdded == 0 && n.getHappiness() > happiest.happiness {
		happiest.path = n.path
		happiest.happiness = n.getHappiness()
	}
}

func (n *seat) getHappiness() int {
	p := strings.Split(n.path, " ")
	h := 0
	for i := 0; i < len(p)-1; i++ {
		p1, p2 := p[i], p[i+1]
		h += hIndexes[p1].rel[p2]
		h += hIndexes[p2].rel[p1]
	}
	p1, p2 := p[0], p[len(p)-1]
	h += hIndexes[p1].rel[p2]
	h += hIndexes[p2].rel[p1]
	return h
}

func parseLine(s string) ([]string, int) {
	s = strings.Replace(s, ".", "", -1)
	a := strings.Split(s, " ")

	p := []string{a[0], a[10]}
	amt := 0

	if a[2] == "lose" {
		amt, _ = strconv.Atoi("-" + a[3])
	} else {
		amt, _ = strconv.Atoi(a[3])
	}

	return p, int(amt)
}
