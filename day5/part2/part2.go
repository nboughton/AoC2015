package main

import (
	"bufio"
	"fmt"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
	"os"
)

var RepeatLetter = pcre.MustCompile(`([a-z]{1})[a-z]{1}\1`, 0)
var RepeatingDoubleLetters = pcre.MustCompile(`([a-z]{2}).*\1`, 0)

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	nice := 0
	for s.Scan() {
		str := s.Text()
		rl := RepeatLetter.MatcherString(str, 0)
		rdl := RepeatingDoubleLetters.MatcherString(str, 0)

		if rl.Matches() && rdl.Matches() {
			nice++
		}
	}

	fmt.Printf("%v strings are nice\n", nice)
}
