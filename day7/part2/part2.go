/*
This one is heavily commented cos it was a pain in the dick to solve
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var isInt = regexp.MustCompile(`^[0-9]+$`)

type wire struct {
	sig      uint16
	op       string
	gateArgs []string
}

type circuit struct {
	plan map[string]*wire
}

type override struct {
	wire   string
	signal uint16
}

func main() {
	// Declare a new circuit
	c := newCircuit()

	// Open the input for reading
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	// read and parse the input
	c.readInput(f)

	// Evaluate wire "a"
	fmt.Printf("Part 1, a: %v\n", c.getSignal("a"))

	// Part 2
	b := c.getSignal("a")
	c = newCircuit()
	f, _ = os.Open(os.Args[1])
	defer f.Close()
	c.readInput(f, &override{wire: "b", signal: b})
	fmt.Printf("Part 2, a: %v\n", c.getSignal("a"))
}

func newCircuit() *circuit {
	c := new(circuit)
	c.plan = make(map[string]*wire)
	return c
}

func (c *circuit) readInput(file *os.File, overrides ...*override) {
	s := bufio.NewScanner(file)
	for s.Scan() {
		// Declare a new wire to be added to circuit plan
		w := new(wire)

		// Get wire id and operation
		atoms := strings.Split(s.Text(), "->")
		op, id := strings.TrimSpace(atoms[0]), strings.TrimSpace(atoms[1])

		// break down operation into components
		atoms = strings.Split(op, " ")
		switch len(atoms) {
		case 1: // Direct Assignment
			if isInt.MatchString(atoms[0]) { // if it's an integer assign it straight to its signal
				v, _ := strconv.ParseInt(atoms[0], 10, 16) // Convert value to uint16
				w.sig = uint16(v)
			} else { // Other wise assign to the arguments so it can be evaluated later
				w.gateArgs = append(w.gateArgs, atoms[0])
			}
		case 2: // NOT operation
			w.op = atoms[0]
			w.gateArgs = append(w.gateArgs, atoms[1])
		case 3: // AND, OR, LSHIFT or RSHIFT operation
			w.op = atoms[1]
			w.gateArgs = append(w.gateArgs, atoms[0], atoms[2])
		}

		// Assign the parsed wire to the circuit plan
		c.plan[id] = w
	}

	// Apply any overrides
	for _, o := range overrides {
		c.setSignal(o.wire, o.signal)
	}
}

// Recursively evaluate the signal for a given wire
func (c *circuit) getSignal(w string) uint16 {
	if !c.hasSignal(w) { // if a wire has no arguments then it has a signal
		// Let's cut to the chase with operation arguments and make them uints that we can pass
		// straight to the operation
		gateArgs := c.plan[w].gateArgs
		opArgs := make([]uint16, len(gateArgs), len(gateArgs))

		// evaluate each argument and see if it has a signal
		for i := 0; i < len(gateArgs); i++ {
			arg := gateArgs[i]           // let's make this easier to read
			if !isInt.MatchString(arg) { // if the arg is not a number it's a wire
				if !c.hasSignal(arg) { // if it has no signal we should get one
					c.setSignal(arg, c.getSignal(arg))
					opArgs[i] = c.plan[arg].sig
				} else if c.hasSignal(arg) { // if it has a signal we should give it to the opArgs
					opArgs[i] = c.plan[arg].sig
				}
			} else { // otherwise it's just a number so give it to the opArgs
				v, _ := strconv.ParseInt(arg, 10, 16)
				opArgs[i] = uint16(v)
			}
		}

		// perform the operation now that all requirements are satisfied and clear the gateArgs
		// to indicate that this wire is complete.
		c.setSignal(w, c.performOperation(w, opArgs))
	}
	return c.plan[w].sig
}

func (c *circuit) setSignal(w string, val uint16) {
	c.plan[w].sig = val
	c.plan[w].gateArgs = nil
}

func (c *circuit) performOperation(w string, args []uint16) uint16 {
	var result uint16
	switch c.plan[w].op {
	case "AND":
		result = args[0] & args[1]
	case "OR":
		result = args[0] | args[1]
	case "LSHIFT":
		result = args[0] << args[1]
	case "RSHIFT":
		result = args[0] >> args[1]
	case "NOT":
		result = ^args[0]
	default:
		result = args[0]
	}
	return result
}

func (c *circuit) hasSignal(w string) bool {
	if len(c.plan[w].gateArgs) == 0 {
		return true
	}
	return false
}
