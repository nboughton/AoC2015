package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Present struct {
	length, width, height, areaX, areaY, areaZ int
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	wrappingPaper := 0
	for s.Scan() {
		p := strings.Split(s.Text(), "x")
		l, _ := strconv.ParseInt(p[0], 10, 32)
		w, _ := strconv.ParseInt(p[1], 10, 32)
		h, _ := strconv.ParseInt(p[2], 10, 32)

		present := NewPresent(int(l), int(w), int(h))
		wrappingPaper += present.SurfaceArea() + present.Slack()
	}

	fmt.Printf("Wrapping Paper: %v\n", wrappingPaper)
}

func NewPresent(l, w, h int) *Present {
	p := &Present{length: l, width: w, height: h}
	p.areaX = p.length * p.width
	p.areaZ = p.width * p.height
	p.areaY = p.height * p.length
	return p
}

func (p *Present) SurfaceArea() int {
	return 2*p.areaX + 2*p.areaZ + 2*p.areaY
}

func (p *Present) Slack() int {
	sides := []int{p.areaX, p.areaZ, p.areaY}
	sort.Ints(sides)
	return sides[0]
}
