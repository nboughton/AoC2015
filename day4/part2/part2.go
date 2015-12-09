package main

import (
	"crypto/md5"
	"fmt"
	"regexp"
)

var Input = "yzbqklnj"
var ValidCoin = regexp.MustCompile(`^0{6}`)

func main() {
	for i := 0; i != -1; i++ {
		coin := md5.Sum([]byte(fmt.Sprintf("%v%v", Input, i)))
		if ValidCoin.MatchString(fmt.Sprintf("%x", coin)) {
			fmt.Printf("Answer: %v\n", i)
			break
		}
	}
}
