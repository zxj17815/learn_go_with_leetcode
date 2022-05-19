package main

import (
	"fmt"
	"math"
)

func main() {
	var tag, other int
	var radix float64
	var n [2]string

	_, _ = fmt.Scanf("%s %s %d %f", &n[0], &n[1], &tag, &radix)
	if tag == 1 {
		other = 2
	} else {
		other = 1
	}

	tagSum := computedSum(radix, n[tag-1])

	var low, high int
	for _, v := range n[other-1] {
		if '0' <= v && v <= '9' {
			low = int(math.Min(float64(low), float64(v-'0')))
		} else {
			low = int(math.Min(float64(low), float64(int(v-'a')+10)))
		}
	}

	low = low + 1
	high = tagSum + 1

	for low <= high {
		mid := (low + high) / 2
		temp := computedSum(float64(mid), n[other-1])
		if temp < 0 {
			high = mid - 1
			continue
		}
		if tagSum < temp {
			high = mid - 1
		} else if tagSum > temp {
			low = mid + 1
		} else {
			fmt.Println(mid)
			return
		}
	}
	fmt.Println("Impossible")
}

func computedSum(radix float64, data string) int {
	var sum float64
	for index, v := range data {
		if '0' <= v && v <= '9' {
			sum += float64(int(v-'0')) * math.Pow(radix, float64(len(data)-index-1))
		} else {
			sum += float64(int(v-'a'+10)) * math.Pow(radix, float64(len(data)-index-1))
		}
	}
	return int(sum)
}
