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

var mostOps = regexp.MustCompile(`(AND|OR|RSHIFT|LSHIFT)`)
var isInt = regexp.MustCompile(`^[0-9]+$`)

type Wire struct {
	sig      uint16
	op       string
	gateArgs []string
}

type Circuit struct {
	plan map[string]*Wire
}

func main() {
	// Declare a new circuit
	c := NewCircuit()

	// Open the input for reading
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	// read and parse the input
	c.ReadInput(f)

	// Evaluate wire "a"
	fmt.Printf("a: %v\n", c.GetSignal("a"))
}

func NewCircuit() *Circuit {
	c := new(Circuit)
	c.plan = make(map[string]*Wire)
	return c
}

func (c *Circuit) ReadInput(file *os.File) {
	s := bufio.NewScanner(file)
	for s.Scan() {
		// Declare a new wire to be added to circuit plan
		w := new(Wire)

		// Get wire id and operation
		l := strings.Split(s.Text(), "->")
		op, id := strings.TrimSpace(l[0]), strings.TrimSpace(l[1])

		// break down operation into components
		l = strings.Split(op, " ")
		if mostOps.MatchString(op) { // If the operation is AND, OR, LSHIFT, RSHIFT
			w.op = l[1]
			w.gateArgs = append(w.gateArgs, l[0], l[2])
		} else if len(l) == 1 { // This is a direct assignment to a wire
			if isInt.MatchString(l[0]) { // if it's an integer assign it straight to its signal
				v, _ := strconv.ParseInt(l[0], 10, 16) // Convert value to uint16
				w.sig = uint16(v)
			} else { // Other wise assign to the arguments so it can be evaluated later
				w.gateArgs = append(w.gateArgs, l[0])
			}
		} else { // Last possible option is that it's a NOT operation
			w.op = l[0]
			w.gateArgs = append(w.gateArgs, l[1])
		}

		// Assign the parsed wire to the circuit plan
		c.plan[id] = w
	}
}

// Recursively evaluate the signal for a given wire
func (c *Circuit) GetSignal(wire string) uint16 {
	if !c.HasSignal(wire) { // if a wire has no arguments then it has a signal
		gateArgs := c.plan[wire].gateArgs
		//fmt.Printf("\n-->%v [%v, %v]\n", wire, gateArgs, c.plan[wire].op)

		// Let's cut to the chase with operation arguments and make them uints that we can pass
		// straight to the operation
		opArgs := make([]uint16, len(gateArgs), len(gateArgs))

		// evaluate each argument and see if it has a signal
		for i := 0; i < len(gateArgs); i++ {
			arg := gateArgs[i] // let's make this easier to read
			//fmt.Printf("?%v\n", arg)

			if !isInt.MatchString(arg) { // if the arg is not a number it's a wire
				if !c.HasSignal(arg) { // if it has no signal we should get one
					c.SetSignal(arg, c.GetSignal(arg))
					opArgs[i] = c.plan[arg].sig

					//fmt.Printf("%v = %v\n", arg, c.plan[arg].sig)
				} else if c.HasSignal(arg) { // if it has a signal we should give it to the opArgs
					opArgs[i] = c.plan[arg].sig
				}
			} else { // otherwise it's just a number so give it to the opArgs
				v, _ := strconv.ParseInt(arg, 10, 16)
				opArgs[i] = uint16(v)
			}
		}

		// Perform the operation now that all requirements are satisfied and clear the gateArgs
		// to indicate that this wire is complete.
		c.plan[wire].sig = c.PerformOperation(wire, opArgs)
		c.plan[wire].gateArgs = nil
	}
	return c.plan[wire].sig
}

func (c *Circuit) SetSignal(wire string, val uint16) {
	c.plan[wire].sig = val
}

func (c *Circuit) PerformOperation(wire string, args []uint16) uint16 {
	var result uint16
	switch c.plan[wire].op {
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

func (c *Circuit) HasSignal(w string) bool {
	if len(c.plan[w].gateArgs) == 0 {
		return true
	}
	return false
}
