package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var containerSets = make(map[int]int)
var cById = make(map[string]int)
var foundSets = make(map[string]int)
var winningSets = make(map[string]int)

var containers []int
var capacity = 150
var total = 0

type node struct {
	path  string
	nodes map[int]*node
}

type set map[int]int

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		v, _ := strconv.Atoi(s.Text())
		containers = append(containers, v)
		containerSets[v]++
	}

	sort.Ints(containers)
	for i, v := range containers {
		si := strconv.Itoa(i)
		sv := strconv.Itoa(v)
		cById[si+":"+sv] = v
	}

	//fmt.Printf("cById %v\n", cById)
	for k, _ := range cById {
		newNode(k).addContainers()
	}

	fmt.Printf("Total permutations: %v\n", len(winningSets))
}

func newNode(path string) *node {
	s := new(node)
	s.nodes = make(map[int]*node)
	s.incrementPath(path)

	return s
}

func (s *node) incrementPath(path string) {
	//fmt.Printf("ip %v\n", path)
	s.path += " " + path + " "
	s.path = strings.TrimSpace(s.path)
}

func (s *node) addContainers() {
	nodesAdded := 0
	for id, container := range cById {
		if !strings.Contains(s.path, id) && s.getTotal()+container <= capacity {
			_, ok := s.nodes[container]
			if !ok {
				s.nodes[container] = newNode(s.path)
			}

			s.nodes[container].incrementPath(id)
			s.nodes[container].addContainers()
			nodesAdded++
		}
	}

	if nodesAdded == 0 && s.getTotal() == capacity {
		s.addSet()
	}
}

func (s *node) addSet() {
	sm := s.path
	if winningSets[sm] != 1 {
		winningSets[sm] = 1
		total++
		fmt.Printf("set %v found: %v\n", total, s.idsToSet())
	}
}

func (s *node) idsToSet() []int {
	st := []int{}
	p := strings.Split(s.path, " ")

	for _, v := range p {
		st = append(st, cById[v])
	}
	return st
}

func (s *node) getTotal() int {
	total := 0
	p := strings.Split(s.path, " ")
	for _, v := range p {
		total += cById[v]
	}
	return total
}
