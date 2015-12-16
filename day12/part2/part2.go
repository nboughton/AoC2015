package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	//"strconv"
	"strings"
)

var isInt = regexp.MustCompile(`([-0-9]+)`)

//var rBlock = regexp.MustCompile(`({[^{]*red[^\}]*})`)
var rBlock = regexp.MustCompile(`({[^{]*red[^}]*})`)

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	//sum := 0
	//r := []string{}
	c := []byte{}
	for s.Scan() {
		/*
			// REGEX solution
			line := s.Text()
			r = rBlock.FindAllString(line, -1)

				d := isInt.FindAllString(line, -1)
				for i := 0; i < len(d); i++ {
					num, _ := strconv.Atoi(d[i])
					sum += num
				}
				for i := 0; i < len(r); i++ {
					num, _ := strconv.Atoi(r[i])
					sum -= num
				}
		*/
		c = removeRedBlocks(s.Bytes())
	}
	/*
		for i := 0; i < len(r); i++ {
			fmt.Printf("sum: %v\n\n", r[i])
		}
	*/
	fmt.Printf("c: %v\n", c)
}

func removeRedBlocks(b []byte) []byte {
	lb := 0                       // index of current left braces
	rb := 0                       // index of current right braces
	prb := []byte{}               // possible red block
	chars := []byte{}             // eventual string
	for i := 0; i < len(b); i++ { // iterate characters
		switch string(b[i]) {
		case "{": // append the index point to lb slice
			lb = i
			fmt.Printf("Open block at i: %v\n", i)
		case "}": // append the index point to rb slice
			rb = i
			fmt.Printf("Close block at i: %v\n", i)
		}
		// if we've opened a block and have the same number of closing braces
		// then check if it contains the string "red"
		if lb != 0 && rb != 0 {
			fmt.Printf("Complete block close at i: %v\n", i)
			prb = b[lb:rb]
			fmt.Printf("Testing prb: %v\n", string(prb))
			break // let's just break at the first one for now so I can see what's going on
			// if not then add it to the chars slice
			if !strings.Contains(string(prb), "red") {
				chars = append(chars, prb...)
				// reset brace slices
				lb, rb = 0, 0
			}
			prb = nil
			// if we're not in a block then just add the character to the chars slice
		} else if lb == 0 {
			chars = append(chars, b[i])
		}
	}
	return chars
}
