package main

import (
	"fmt"
	"sort"
)

func main() {
	poly := make(map[int]float64)
	_map := make(map[int]float64)
	var k int
	_, _ = fmt.Scanf("%d", &k)
	for i := 0; i < k; i++ {
		var key int
		var value float64
		_, _ = fmt.Scanf("%d", &key)
		_, _ = fmt.Scanf("%f", &value)
		poly[key] = value
	}

	_, _ = fmt.Scanf("%d", &k)
	for i := 0; i < k; i++ {
		var key int
		var value float64
		_, _ = fmt.Scanf("%d", &key)
		_, _ = fmt.Scanf("%f", &value)
		for index, v := range poly {
			tempK := index + key
			tempValue := v * value
			val, exist := _map[tempK]
			if exist {
				_map[tempK] = val + tempValue
			} else {
				_map[tempK] = tempValue
			}
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
