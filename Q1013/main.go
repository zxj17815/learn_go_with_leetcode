package main

import (
	"fmt"
)

var n, m, k int
var contentMapByCity [][]int
var path map[int]bool
var count int

func main() {

	_, _ = fmt.Scanf("%d %d %d", &n, &m, &k)
	contentMapByCity = make([][]int, n)
	for i := 0; i < m; i++ {
		var city1, city2 int
		_, _ = fmt.Scanf("%d %d", &city1, &city2)
		contentMapByCity[city1-1] = append(contentMapByCity[city1-1], city2-1)
		contentMapByCity[city2-1] = append(contentMapByCity[city2-1], city1-1)
	}
	for i := 0; i < k; i++ {
		var city int
		path = make(map[int]bool)
		_, _ = fmt.Scanf("%d", &city)
		path[city-1] = true
		count = 0
		for index := range contentMapByCity {
			count += deep(index)
		}
		fmt.Println(count - 1)
	}
}

func deep(i int) int {
	if path[i] {
		return 0
	} else {
		path[i] = true
	}
	for _, v := range contentMapByCity[i] {
		deep(v)
	}
	return 1
}
