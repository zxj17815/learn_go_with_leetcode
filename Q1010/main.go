package main

import (
	"fmt"
	"strconv"
)

func main() {
	var radix string
	var tag int
	data := make([]string, 2)
	a_z := "abcdefghijklmnopqrstuvwxyz"
	radixList := make(map[string]int)
	for i := 0; i < 10; i++ {
		radixList[strconv.Itoa(i)] = i
	}
	for i := 0; i < len(a_z); i++ {
		radixList[string(a_z[i])] = i + 10
	}

	_, _ = fmt.Scanf("%s %s %d %s", &data[0], &data[1], &tag, &radix)

	find := false

	var val int
	for i := len(data[tag-1]) - 1; i > 0; i-- {
		var += data[tag-1][i]
	}
	for i := 0; i < len(radixList); i++ {
		if radix == radixList[i] {
			tag = i
		}
	}
}
