package main

import "fmt"

func main() {
	var n, d int
	for {
		_, _ = fmt.Scanf("%d", &n)
		if n <= 1 {
			break
		}
		_, _ = fmt.Scanf("%d", &d)
		if !isPrime(n) {
			fmt.Println("No")
			continue
		}

		length := 0
		number := [20]int{0}

		for {
			number[length] = n % d
			length++
			n = n / d
			if n <= 0 {
				break
			}
		}
		for i := 0; i < length; i++ {
			n = n*d + number[i]
		}
		if isPrime(n) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
func isPrime(i int) bool {
	for j := 2; j < i; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}
