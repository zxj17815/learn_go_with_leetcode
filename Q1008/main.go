package main

import "fmt"

func main() {
	var n, floor, temp, time int
	_, _ = fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &temp)
		if floor < temp {
			time += 6*(temp-floor) + 5
		} else {
			time += 4*(floor-temp) + 5
		}
		floor = temp
	}
	fmt.Println(time)
}
