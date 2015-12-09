package main

import (
	"bufio"
	"fmt"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
	"os"
)

var Vowel = pcre.MustCompile(`[aeiou]{1}.*[aeiou]{1}.*[aeiou]{1}`, 0)
var RepeatLetter = pcre.MustCompile(`([a-z]{1})\1`, 0)
var CannotContain = pcre.MustCompile(`(ab|cd|pq|xy)`, 0)

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	nice := 0
	for s.Scan() {
		str := s.Text()
		v := Vowel.MatcherString(str, 0)
		rl := RepeatLetter.MatcherString(str, 0)
		cc := CannotContain.MatcherString(str, 0)

		if v.Matches() && rl.Matches() && !cc.Matches() {
			nice++
		}
	}

	fmt.Printf("%v strings are nice\n", nice)
}
