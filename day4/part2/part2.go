package main

import (
	"crypto/md5"
	"fmt"
	"regexp"
)

var input = "yzbqklnj"
var validCoin = regexp.MustCompile(`^0{6}`)

func main() {
	for i := 0; i != -1; i++ {
		coin := md5.Sum([]byte(fmt.Sprintf("%v%v", input, i)))
		if validCoin.MatchString(fmt.Sprintf("%x", coin)) {
			fmt.Printf("Answer: %v\n", i)
			break
		}
	}
}
