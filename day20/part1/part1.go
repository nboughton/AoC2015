package main

import (
	"fmt"
)

func main() {
	input := 34000000
	var house, presents int

	// houses 2,4,8,16 and all primes etc = n-1 * 2 + 10
	// Increment houses by 60 as multiples of 60 produce the
	// biggest step change in present numbers
	for house = 0; house != -1; house += 60 {
		presents = getsPresents(house)
		if presents >= input {
			break
		}
	}
}

func getsPresents(house int) int {
	presents := 0
	for elf := 1; elf <= house; elf++ {
		if house%elf == 0 {
			presents += 10 * elf
		}
	}
	fmt.Printf("house %v gets %v presents\n", house, presents)
	return presents
}
