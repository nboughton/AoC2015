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
		fmt.Printf("Node 1: %v\n", a)
		n[a] = newNode(a, a)
		n[a].addNodes(places)
	}

	fmt.Printf("places %v\n", n)
}

func newNode(name, path string) *node {
	n := new(node)
	n.connects = make(map[string]*node)
	n.name = name
	n.path += path + " "

	fmt.Printf("Node: %v\n", n)
	return n
}

func (n *node) addNodes(places map[string]*place) {
	fmt.Printf("Name: %v\n", n.name)
	for place := range places {
		if !strings.Contains(n.path, place) {
			_, ok := n.connects[place]
			if !ok {
				fmt.Printf("Connect new node: %v\n", place)
				n.connects[place] = newNode(place, place)
			}

			n.path += place + " "
			n.connects[place].addNodes(places)
			fmt.Printf("Path: %v\n\n", n.path)
		}
	}
}

func parseLine(s string) ([]string, int) {
	atoms := strings.Split(s, " ")
	places := []string{atoms[0], atoms[2]}
	distance, _ := strconv.ParseInt(atoms[4], 10, 32)

	return places, int(distance)
}
