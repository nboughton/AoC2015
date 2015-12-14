package main

import (
	"fmt"
	"regexp"
)

var hex = regexp.MustCompile(`\\x[a-z0-9]{2}`)
var dbs = regexp.MustCompile(`\\[\\"]{1}`)

func main() {
	str := `"asdf\"\xd3\x63\\"`
	fmt.Println(dbs.FindAllString(str, -1))
}
