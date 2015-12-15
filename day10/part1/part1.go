package main

import (
	"fmt"
	"strconv"
)

var input = []byte("1113222113")

func main() {
	output := input
	for i := 0; i < 40; i++ {
		output = lookAndSay(output)
		fmt.Printf("i:%v, string length: %v\n", i, len(output))
	}
}

func lookAndSay(s []byte) []byte {
	o := []byte{}
	c := 1
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if s[i] == s[i+1] {
				c++
			} else {
				b := []byte(strconv.FormatInt(int64(c), 10))
				o = append(o, b[0], s[i])
				c = 1
			}
		} else {
			if s[i] == s[i-1] {
				b := []byte(strconv.FormatInt(int64(c), 10))
				o = append(o, b[0], s[i])
			} else {
				b := []byte(strconv.FormatInt(int64(c), 10))
				c = 1
				o = append(o, b[0], s[i])
			}
		}
	}
	return o
}
