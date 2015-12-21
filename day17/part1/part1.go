package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
)

var containerSets = make(map[int]int)
var uniqueSets []*set
var foundSets []set

var containers []int
var capacity = 150
var total = 0

type node struct {
	path  *set
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
	for _, c := range containers {
		newNode(newSet(c)).addContainers()
	}

	fmt.Printf("Total permutations: %v\n", len(uniqueSets))
}

func newNode(path *set) *node {
	s := new(node)
	s.nodes = make(map[int]*node)
	s.path = newSet()
	s.incrementPath(path)

	return s
}

func newSet(p ...int) *set {
	s := make(set)

	for i := 0; i < len(p); i++ {
		s[p[i]]++
	}

	return &s
}

func (s *node) incrementPath(path *set) {
	for k, v := range *path {
		(*s.path)[k] += v
	}
}

func (s *node) addContainers() {
	for _, container := range containers {
		if s.checkSet(container) && s.getTotal()+container <= capacity {
			_, ok := s.nodes[container]
			if !ok {
				s.nodes[container] = newNode(s.path)
			}

			s.nodes[container].incrementPath(newSet(container))
			if s.nodes[container].checkSet() {
				//fmt.Printf("p: %v\n", s.path)
				s.nodes[container].addContainers()
			} else {
				delete(s.nodes, container)
			}
		}
	}

	if s.checkSet() && s.getTotal() == capacity {
		s.addSet()
	}
}

func (s *node) addSet() {
	found := false

	for _, v := range uniqueSets {
		if reflect.DeepEqual(s.path, v) {
			found = true
		}
	}

	if !found {
		total++
		fmt.Printf("set %v found: %v\n", total, s.path)
		uniqueSets = append(uniqueSets, s.path)
	}
}

func (s *node) checkSet(t ...int) bool {
	control := make(map[int]int)
	if len(t) > 0 {
		control[t[0]]++
	}

	for k, v := range *s.path {
		control[k] = v
	}

	for k, _ := range control {
		if control[k] > containerSets[k] {
			return false
		}
	}

	return true
}

func (s *node) getTotal() int {
	total := 0

	for k, v := range *s.path {
		total += k * v
	}
	return total
}
