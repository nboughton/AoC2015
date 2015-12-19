package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var isInt = regexp.MustCompile(`([-0-9]+)`)

type ingredient struct {
	name                                             string
	capacity, durability, flavour, texture, calories int
}

type cookie struct {
	capacity, durability, flavour, texture, calories, score int
}

var stock []*ingredient
var winner int

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		stock = append(stock, parseLine(s.Text()))
	}

	m := make([]int, len(stock))
	makeCookies(100, 0, m)

	fmt.Printf("The winning cookie has a score of %v\n", winner)
}

func makeCookies(ts, i int, m []int) {
	ot := 0
	for j := 0; j < i; j++ {
		ot += m[j]
	}

	for m[i] = 1; m[i]+ot <= ts; m[i]++ {
		if i != len(m)-1 {
			makeCookies(ts, i+1, m)
		}
		if getTotal(m) == 100 {
			c := getScore(m)
			if c.score > winner && c.calories == 500 {
				winner = c.score
			}
		}
	}
}

func getScore(m []int) *cookie {
	c := new(cookie)
	for i := 0; i < len(m); i++ {
		c.capacity += stock[i].capacity * m[i]
		c.durability += stock[i].durability * m[i]
		c.flavour += stock[i].flavour * m[i]
		c.texture += stock[i].texture * m[i]
		c.calories += stock[i].calories * m[i]
	}

	if c.capacity < 0 {
		c.capacity = 0
	}
	if c.durability < 0 {
		c.durability = 0
	}
	if c.flavour < 0 {
		c.flavour = 0
	}
	if c.texture < 0 {
		c.texture = 0
	}
	if c.calories < 0 {
		c.calories = 0
	}

	c.score = c.capacity * c.durability * c.flavour * c.texture
	return c
}

func getTotal(m []int) int {
	t := 0
	for i := 0; i < len(m); i++ {
		t += m[i]
	}
	return t
}

func parseLine(s string) *ingredient {
	a := strings.Split(s, ":")
	name := a[0]

	a = isInt.FindAllString(a[1], -1)

	ca, _ := strconv.Atoi(a[0])
	du, _ := strconv.Atoi(a[1])
	fl, _ := strconv.Atoi(a[2])
	te, _ := strconv.Atoi(a[3])
	cl, _ := strconv.Atoi(a[4])

	return &ingredient{
		name:       name,
		capacity:   ca,
		durability: du,
		flavour:    fl,
		texture:    te,
		calories:   cl,
	}
}
