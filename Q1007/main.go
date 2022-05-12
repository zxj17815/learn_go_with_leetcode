package main

import "fmt"

func main() {

	var k int
	var negative int
	max := ^int(^uint(0) >> 1)
	_, _ = fmt.Scanf("%d", &k)
	sub := make([]int, k)
	for i := 0; i < k; i++ {
		_, _ = fmt.Scanf("%d", &sub[i])
		if sub[i] < 0 {
			negative++
		}
	}
	if negative == k {
		fmt.Println(0, sub[0], sub[k-1])
		return
	}
	if k == 1 {
		fmt.Println(sub[0], sub[0], sub[0])
		return
	}
	var sum, start, end int
	for i := 0; i < k-1; i++ {
		for j := i; j < k; j++ {
			sum = 0
			for _, v := range sub[i : j+1] {
				sum += v
			}
			if sum > max {
				max = sum
				start = i
				end = j
			}

		}
	}

	fmt.Println(max, sub[start], sub[end])
}
