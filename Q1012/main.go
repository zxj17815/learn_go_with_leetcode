package main

import (
	"fmt"
	"sort"
)

func main() {
	scoreMap := map[string][]int{}
	var n, m int
	var courses [4][]int
	rankMap := map[int]string{0: "A", 1: "C", 2: "M", 3: "E"}

	_, _ = fmt.Scanf("%d %d", &n, &m)

	for i := 0; i < n; i++ {
		var name string
		var score [3]int
		_, _ = fmt.Scanf("%s %d %d %d", &name, &score[0], &score[1], &score[2])
		a := (score[0] + score[1] + score[2]) / 3
		courses[0] = append(courses[0], a)
		courses[1] = append(courses[1], score[0])
		courses[2] = append(courses[2], score[1])
		courses[3] = append(courses[3], score[2])
		scoreMap[name] = append(scoreMap[name], a, score[0], score[1], score[2])
	}
	for k := range courses {
		sort.Sort(sort.Reverse(sort.IntSlice(courses[k])))
	}

	for i := 0; i < m; i++ {
		var tempId string
		_, _ = fmt.Scanf("%s", &tempId)
		if _, ok := scoreMap[tempId]; ok {
			rank := n + 1
			var courser string
			for k, v := range scoreMap[tempId] {
				tempRank := find(v, courses[k])
				if rank > tempRank {
					rank = tempRank
					courser = rankMap[k]
				}
			}
			fmt.Printf("%d %s\n", rank, courser)
		} else {
			fmt.Println("N/A")
		}
	}
}

func find(i int, list []int) int {
	for k, v := range list {
		if v == i {
			return k + 1
		}
	}
	return -1
}
