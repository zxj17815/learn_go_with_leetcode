package main

import "fmt"

func main() {
	var n, m, k, q int
	_, _ = fmt.Scanf("%d %d %d %d", &n, &m, &k, &q)
	line := make([][]int, n)
	lineIndex := make([][]int, n)
	init := n * m
	temp := 0
	for i := 0; i < init; i++ {
		if temp == n-1 {
			temp = 0
		} else {
			temp++
		}
		var time int
		_, _ = fmt.Scanf("%d", &time)
		lineIndex[temp] = append(lineIndex[temp], len(line[temp]))
		if len(line[temp]) > 0 {
			line[temp] = append(line[temp], line[temp][i-1]+time)
		} else {
			line[temp] = append(line[temp], time)
		}

	}

	for i := 0; i < k-init; i++ {
		var time, addLine int
		minTime := 999999999
		_, _ = fmt.Scanf("%d", &time)
		for j := 0; j < n; j++ {
			if line[j][len(line[j])-1] < minTime {
				minTime = line[j][len(line[j])-1]
				addLine = j
			}
		}
		lineIndex[addLine] = append(lineIndex[addLine], line[addLine][len(line[addLine])-1]+time)
	}
	//todo 输出数据 八点开始加上分钟数量
}
