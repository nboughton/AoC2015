package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type place struct {
	name     string
	connects map[string]int
}

type node struct {
	name, path string
	connects   map[string]*node
	distance   int
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	places := make(map[string]*place)

	for s.Scan() {
		n, distance := parseLine(s.Text())

		for i := 0; i < len(n); i++ {
			_, ok := places[n[i]]
			if !ok {
				places[n[i]] = new(place)
				places[n[i]].connects = make(map[string]int)
			}
		}
		places[n[0]].connects[n[1]] = distance
		places[n[1]].connects[n[0]] = distance
	}

	//distance := 0
	n := make(map[string]*node)
	for a := range places {
		n[a] = newNode(a, a)
		n[a].addNodes(places)
	}

	fmt.Printf("places %v\n", n)
}

func newNode(name, path string) *node {
	n := new(node)
	n.connects = make(map[string]*node)
	n.name = name
	n.path += path

	return n
}

func (n *node) incrementPath(path string) {
	n.path += path
}

func (n *node) addNodes(places map[string]*place) {
	for place := range places {
		if !strings.Contains(n.path, place) {
			_, ok := n.connects[place]
			if !ok {
				n.connects[place] = newNode(place, n.path)
			}

			n.incrementPath(place)
			n.connects[place].incrementPath(place)

			n.connects[place].addNodes(places)
		}
	}
}

func parseLine(s string) ([]string, int) {
	atoms := strings.Split(s, " ")
	places := []string{atoms[0], atoms[2]}
	distance, _ := strconv.ParseInt(atoms[4], 10, 32)

	return places, int(distance)
}
