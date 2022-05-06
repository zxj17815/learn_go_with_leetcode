package main

import (
	"fmt"
	"sort"
)

func main() {
	_map := make(map[int]float64)
	var k int
	_, _ = fmt.Scanf("%d", &k)
	for i := 0; i < k; i++ {
		var key int
		var value float64
		_, _ = fmt.Scanf("%d", &key)
		_, _ = fmt.Scanf("%f", &value)
		_map[key] = value
	}

	_, _ = fmt.Scanf("%d", &k)
	for i := 0; i < k; i++ {
		var key int
		var value float64
		_, _ = fmt.Scanf("%d", &key)
		_, _ = fmt.Scanf("%f", &value)
		val, exist := _map[key]
		if exist {
			_map[key] = val + value
		} else {
			_map[key] = value
		}
	}

	sum := 0
	var keyArr []int
	for key, val := range _map {
		keyArr = append(keyArr, key)
		if val != 0 {
			sum++
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keyArr)))

	fmt.Printf("%d", sum)

	for _, key := range keyArr {
		val := _map[key]
		if val != 0 {
			fmt.Printf(" %d %.1f", key, val)
		}
	}

}
