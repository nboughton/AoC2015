package main

import (
	"bufio"
	"fmt"
	"github.com/ntns/goitertools/itertools"
	"os"
	"strconv"
)

var cVal []int
var total int

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		v, _ := strconv.Atoi(s.Text())
		cVal = append(cVal, v)
	}

	minChk := false
	for i := 2; i < len(cVal)-1; i++ {
		com := itertools.Combinations(cVal, i)
		for _, l := range com {
			if getTotal(l) == 150 {
				minChk = true
				total++
				fmt.Printf("c:%v\n", l)
			}
		}
		if minChk {
			break
		}
	}

	fmt.Printf("Total sets: %v\n", total)
}

func getTotal(num []int) int {
	t := 0
	for _, v := range num {
		t += v
	}
	return t
}
