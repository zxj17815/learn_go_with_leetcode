package main

import (
	"fmt"
	"strconv"
)

func main() {
	spell := [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var s string
	var sum int

	_, _ = fmt.Scanf("%s", &s)
	for _, v := range s {
		a, _ := strconv.Atoi(string(v))
		sum += a
	}

	s = strconv.Itoa(sum)

	index, _ := strconv.Atoi(string(s[0]))
	fmt.Printf("%s", spell[index])

	if sum < 10 {
		return
	} else {
		for _, v := range s[1:] {
			index, _ = strconv.Atoi(string(v))
			fmt.Printf(" %s", spell[index])
		}
	}
}
