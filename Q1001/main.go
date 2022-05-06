package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	for {
		n, _ := fmt.Scanf("%d %d", &a, &b)
		if n != 2 {
			break
		}
		if a+b < 0 {
			fmt.Println("-" + separatedByCommas(strconv.Itoa(-1*(a+b))))
		} else {
			fmt.Println(separatedByCommas(strconv.Itoa(a + b)))
		}
	}
}

func separatedByCommas(s string) string {
	if len(s) < 4 {
		return s
	} else {
		return separatedByCommas(s[:len(s)-3]) + "," + s[len(s)-3:]
	}
}
