package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = "1113222113"

func main() {
	output := input
	for i := 0; i < 40; i++ {
		output = lookAndSay(output)
		fmt.Printf("i:%v, string length: %v\n", i, len(output))
	}
}

func lookAndSay(s string) string {
	o := ""
	a := strings.Split(s, "")
	c := 1
	for i := 0; i < len(a); i++ {
		if i+1 < len(a) {
			if a[i] == a[i+1] {
				c++
			} else {
				o += strconv.FormatInt(int64(c), 10) + a[i]
				c = 1
			}
		} else {
			if a[i] == a[i-1] {
				o += strconv.FormatInt(int64(c), 10) + a[i]
			} else {
				c = 1
				o += strconv.FormatInt(int64(c), 10) + a[i]
			}
		}
	}
	return o
}
