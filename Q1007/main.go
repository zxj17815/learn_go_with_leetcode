package main

import "fmt"

func main() {

	var k, max, sum, start, end, index int
	_, _ = fmt.Scanf("%d", &k)
	sub := make([]int, k)
	max = -1
	end = k - 1
	for i := 0; i < k; i++ {
		_, _ = fmt.Scanf("%d", &sub[i])
		sum = sum + sub[i]
		if sum < 0 {
			sum = 0
			index = i + 1
		} else if sum > max {
			max = sum
			start = index
			end = i
		}
	}
	if max < 0 {
		max = 0
	}
	fmt.Println(max, sub[start], sub[end])
}
