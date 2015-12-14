package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type place struct {
	name string
	to   map[string]int
}

// places is a global reference for nodes and distances between nodes
var places = make(map[string]*place)

type node struct {
	name, path string
	connects   map[string]*node
}

type route struct {
	distance int
	path     string
}

var longest = &route{distance: 0}
var shortest = &route{distance: 1000}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		n, distance := parseLine(s.Text())

		for i := 0; i < len(n); i++ {
			_, ok := places[n[i]]
			if !ok {
				places[n[i]] = new(place)
				places[n[i]].to = make(map[string]int)
			}
		}
		places[n[0]].to[n[1]] = distance
		places[n[1]].to[n[0]] = distance
	}

	for a := range places {
		n := newNode(a, a)
		n.addNodes()
	}

	fmt.Printf("Shortest route: %v, distance: %v\n", shortest.path, shortest.distance)
	fmt.Printf("Longest route: %v, distance: %v\n", longest.path, longest.distance)
}

func newNode(name, path string) *node {
	n := new(node)
	n.connects = make(map[string]*node)
	n.name = name
	n.incrementPath(path)

	return n
}

func (n *node) incrementPath(path string) {
	n.path += path + " "
	n.path = strings.Replace(n.path, "  ", " ", -1)
}

func (n *node) addNodes() {
	nodesAdded := 0
	for place := range places {
		if !strings.Contains(n.path, place) {
			_, ok := n.connects[place]
			if !ok {
				n.connects[place] = newNode(place, n.path)
			}

			n.connects[place].incrementPath(place)
			n.connects[place].addNodes()
			nodesAdded++
		}
	}

	if nodesAdded == 0 {
		if n.getDistance() < shortest.distance {
			shortest.path = n.path
			shortest.distance = n.getDistance()
		} else if n.getDistance() > longest.distance {
			longest.path = n.path
			longest.distance = n.getDistance()
		}
	}
}

func (n *node) getDistance() int {
	p := strings.Split(n.path, " ")
	d := 0
	for i := 0; i < len(p); i++ {
		if i+1 < len(p) {
			d += places[p[i]].to[p[i+1]]
		}
	}
	return d
}

func parseLine(s string) ([]string, int) {
	atoms := strings.Split(s, " ")
	places := []string{atoms[0], atoms[2]}
	distance, _ := strconv.ParseInt(atoms[4], 10, 32)

	return places, int(distance)
}
