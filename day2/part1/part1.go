package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type present struct {
	length, width, height, areaX, areaY, areaZ int
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	wrp := 0
	for s.Scan() {
		a := strings.Split(s.Text(), "x")
		l, _ := strconv.ParseInt(a[0], 10, 32)
		w, _ := strconv.ParseInt(a[1], 10, 32)
		h, _ := strconv.ParseInt(a[2], 10, 32)

		p := newPresent(int(l), int(w), int(h))
		wrp += p.surfaceArea() + p.slack()
	}

	fmt.Printf("Wrapping Paper: %v\n", wrp)
}

func newPresent(l, w, h int) *present {
	p := &present{length: l, width: w, height: h}
	p.areaX = p.length * p.width
	p.areaZ = p.width * p.height
	p.areaY = p.height * p.length
	return p
}

func (p *present) surfaceArea() int {
	return 2*p.areaX + 2*p.areaZ + 2*p.areaY
}

func (p *present) slack() int {
	sides := []int{p.areaX, p.areaZ, p.areaY}
	sort.Ints(sides)
	return sides[0]
}
