package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var total float64

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	var j interface{} // stuff goes in here
	for s.Scan() {
		json.Unmarshal(s.Bytes(), &j)
	}
	checkArr(j.([]interface{}))
	fmt.Printf("Total: %v\n", total)
}

func checkArr(m []interface{}) {
	for _, v := range m {
		switch vv := v.(type) {
		case []interface{}:
			checkArr(vv)
		case map[string]interface{}:
			checkObj(vv)
		case float64:
			total += vv
		}
	}
}

func checkObj(m map[string]interface{}) {
	red := false

	children := []interface{}{}

	for _, v := range m {
		switch vv := v.(type) {
		case string:
			if vv == "red" {
				red = true
				break
			}
		case map[string]interface{}: // new obj
			children = append(children, vv)
		case []interface{}: // array
			children = append(children, vv)
		}
	}

	if !red {
		for _, v := range m {
			switch vv := v.(type) {
			case float64:
				total += vv
			}
		}
		checkArr(children)
	}
}
