package main

import (
	"fmt"
	"reflect"
	"regexp"
)

var input = []byte("hxbxwxba")

var abc = []byte("abcdefghijklmnopqrstuvwxyz")
var atoi = make(map[byte]int)
var itoa = make(map[int]byte)

var cannotContain = regexp.MustCompile(`(i|o|l)`)

func main() {
	// Create two maps of letter <-> integer relationships
	for i, l := range abc {
		atoi[l] = i + 1
		itoa[i+1] = l
	}

	_ = getNextPassword(input)
}

func getNextPassword(s []byte) []byte {
	for true {
		s = incrementPassword(s, len(s)-1)
		//fmt.Printf("Testing %v\n", string(s))
		if secDoubleLetters(s) && secThreeAscending(s) && !cannotContain.Match(s) {
			fmt.Printf("Found %v\n", string(s))
			break
		}
	}
	return s
}

func secDoubleLetters(s []byte) bool {
	d := []int{}
	for i := 0; i < len(s)-1; i++ {
		t := []int{atoi[s[i]], atoi[s[i+1]]}
		if t[0] == t[1] {
			d = append(d, t[0])
			if len(d) == 2 && d[0] != d[1] {
				return true
			}
		}
	}
	return false
}

func secThreeAscending(s []byte) bool {
	for i := 0; i < len(s)-2; i++ {
		t1 := []int{atoi[s[i]], atoi[s[i+1]], atoi[s[i+2]]} // a, b, c
		t2 := []int{}
		for j := 0; j < len(abc)-2; j++ {
			t2 = []int{atoi[abc[j]], atoi[abc[j+1]], atoi[abc[j+2]]} // a, b, c -> b, c, d
			if reflect.DeepEqual(t1, t2) {
				return true
			}
		}
	}
	return false
}

// increments the letter and returns true/false if it wraps
func incrementLetter(l byte) (byte, bool) {
	i := atoi[l]
	w := false
	if i+1 > len(abc) {
		w = true
		return itoa[1], w
	}
	return itoa[i+1], w
}

func incrementPassword(b []byte, i int) []byte {
	w := false
	b[i], w = incrementLetter(b[i])
	if w && i-1 >= 0 {
		incrementPassword(b, i-1)
	}
	return b
}
