package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var MatchInt = regexp.MustCompile(`^[0-9]+$`)

type Circuit map[string]chan uint16

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	circuit := NewCircuit()

	for s.Scan() {
		line := s.Text()
		go func() {
			atoms := strings.Split(line, "->")
			gate, wire := strings.TrimSpace(atoms[0]), strings.TrimSpace(atoms[1])

			circuit.CreateWire(wire)

			atoms = strings.Split(gate, " ")
			switch len(atoms) {
			case 1: // Direct assignment
				circuit[wire] <- circuit.PatchSignal(atoms[0])
			case 2: // NOT operator
				circuit[wire] <- ^circuit.PatchSignal(atoms[1])
			case 3: // AND|OR|LSHIFT|RSHIFT operators
				l := circuit.PatchSignal(atoms[0])
				r := circuit.PatchSignal(atoms[2])
				switch atoms[1] {
				case "AND":
					circuit[wire] <- l & r
				case "OR":
					circuit[wire] <- l | r
				case "LSHIFT":
					circuit[wire] <- l << r
				case "RSHIFT":
					circuit[wire] <- l >> r
				}
			}
			close(circuit[wire])
		}()
	}

	fmt.Printf("Stupid sexy flanders: %v\n", <-circuit["a"])
}

func NewCircuit() Circuit {
	c := make(Circuit)
	return c
}

func (c Circuit) CreateWire(w string) {
	_, ok := c[w]
	if !ok {
		fmt.Printf("Creating wire: %v\n", w)
		c[w] = make(chan uint16, 1)
	}
}

func (c Circuit) PatchSignal(s string) uint16 {
	var v uint16
	if !isInt(s) {
		c.CreateWire(s)
		v = <-c[s]
	} else {
		t, _ := strconv.ParseInt(s, 10, 16)
		v = uint16(t)
	}
	return v
}

func isInt(s string) bool {
	if MatchInt.MatchString(s) {
		return true
	}
	return false
}
