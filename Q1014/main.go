package main

import (
	"fmt"
	"time"
)

type window struct {
	first int
	index int
	queue []int
}

func main() {
	var n, m, k, q int
	_, _ = fmt.Scanf("%d %d %d %d", &n, &m, &k, &q)
	windows := make([]window, n)
	windowIndex := make([][]int, k)
	init := n * m
	temp := 0
	if init >= k {
		init = k
	}
	for i := 0; i < init; i++ {
		var costTime int
		_, _ = fmt.Scanf("%d", &costTime)
		windowIndex[i] = append(windowIndex[i], temp, len(windows[temp].queue))
		if len(windows[temp].queue) <= 0 {
			windows[temp].first = costTime
			windows[temp].index = 0
		}
		windows[temp].queue = append(windows[temp].queue, costTime)
		if temp == n-1 {
			temp = 0
		} else {
			temp++
		}
	}
	if n*m < k {
		for i := 0; i < k-init; i++ {
			var t, addWindow int
			minTime := 999999999
			_, _ = fmt.Scanf("%d", &t)
			for j := 0; j < n; j++ {
				if windows[j].first < minTime {
					minTime = windows[j].first
					addWindow = j
				}
			}
			windowIndex[init+i] = append(windowIndex[init+i], addWindow, len(windows[addWindow].queue))
			windows[addWindow].queue = append(windows[addWindow].queue, t)
			windows[addWindow].index++
			windows[addWindow].first = windows[addWindow].queue[windows[addWindow].index]
			for j := 0; j < n; j++ {
				if j != addWindow {
					if windows[j].first-minTime == 0 {
						windows[j].index++
						windows[j].first = windows[j].queue[windows[j].index]
					} else {
						windows[j].first = windows[j].first - minTime
					}
				}
			}
		}
	}
	//todo 输出数据 八点开始加上分钟数量
	for i := 0; i < q; i++ {
		var query int
		_, _ = fmt.Scanf("%d", &query)
		c := count(windows[windowIndex[query-1][0]].queue[:windowIndex[query-1][1]+1])
		if count(windows[windowIndex[query-1][0]].queue[:windowIndex[query-1][1]]) < 9*60 {
			t, _ := time.Parse("15:04", "8:00")
			minute, _ := time.ParseDuration("1m")
			fmt.Println(t.Add(minute * time.Duration(c)).Format("15:04"))
		} else {
			fmt.Println("Sorry")
		}
	}
}

func count(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}
