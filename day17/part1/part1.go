package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		v, _ := strconv.Atoi(s.Text())
		containers = append(containers, v)
	}

	sort.Ints(containers)
	for i, v := range containers {
		si := strconv.Itoa(i)
		sv := strconv.Itoa(v)
		cById[si+":"+sv] = v
	}

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
	s.path += " " + path + " "
	s.path = strings.TrimSpace(s.path)
}

func (s *node) addContainers() {
	for id, container := range cById {
		if !strings.Contains(s.path, id) && s.getTotal()+container < capacity {
			_, ok := s.nodes[container]
			if !ok {
				s.nodes[container] = newNode(s.path)
			}

			s.nodes[container].incrementPath(id)
			if s.nodes[container].validSet() {
				s.nodes[container].addContainers()
			} else {
				delete(s.nodes, container)
			}
		}
	}
	if s.validSet() && s.getTotal() == capacity {
		s.addSet()
	}
}

func (s *node) validSet() bool {
	p := strings.Split(s.path, " ")

	control := make(map[string]int)
	for _, v := range p {
		control[v]++
		if control[v] > 1 {
			return false
		}
	}

	return true
}

func (s *node) addSet() {
	sm := ""
	for _, v := range s.idsToSet() {
		sm += strconv.Itoa(v) + " "
	}
	sm = strings.TrimSpace(sm)

	if winningSets[sm] != 1 {
		winningSets[sm] = 1
		total++
		fmt.Printf("set %v found: %v\n", total, sm)
	}
}

func (s *node) idsToSet() []int {
	st := []int{}
	p := strings.Split(s.path, " ")

	for _, v := range p {
		st = append(st, cById[v])
	}
	sort.Ints(st)
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
