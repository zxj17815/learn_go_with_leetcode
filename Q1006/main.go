package main

import (
	"fmt"
	"time"
)

func main() {
	var m int
	var firstId, lastId string
	firstIn, _ := time.Parse("15:04:05", "23:59:59")
	lastOut, _ := time.Parse("15:04:05", "00:00:00")
	_, _ = fmt.Scanf("%d", &m)
	for i := 0; i < m; i++ {
		var id string
		var temp string
		_, _ = fmt.Scanf("%s", &id)

		_, _ = fmt.Scanf("%s", &temp)
		in, _ := time.Parse("15:04:05", temp)
		_, _ = fmt.Scanf("%s", &temp)
		out, _ := time.Parse("15:04:05", temp)

		if in.Before(firstIn) {
			firstIn = in
			firstId = id
		}
		if out.After(lastOut) {
			lastId = id
			lastOut = out
		}
	}
	fmt.Println(firstId, lastId)
}
