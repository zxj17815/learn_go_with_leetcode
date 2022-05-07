package main

import "fmt"

var cities, roads, cur, emer int
var cityTeams []int
var cityVisited []bool
var roadInfo [][][]int
var roadLen, minRoadLen, teams, maxTeams, minRoads int

func main() {
	minRoadLen = int(^uint(0) >> 1)
	_, _ = fmt.Scanf("%d", &cities)
	_, _ = fmt.Scanf("%d", &roads)
	_, _ = fmt.Scanf("%d", &cur)
	_, _ = fmt.Scanf("%d", &emer)

	for i := 0; i < cities; i++ {
		var team int
		_, _ = fmt.Scanf("%d", &team)
		cityTeams = append(cityTeams, team)
	}

	roadInfo = make([][][]int, roads)
	cityVisited = make([]bool, cities)

	for i := 0; i < roads; i++ {
		var city1, city2, L int
		_, _ = fmt.Scanf("%d", &city1)
		_, _ = fmt.Scanf("%d", &city2)
		_, _ = fmt.Scanf("%d", &L)
		roadInfo[city1] = append(roadInfo[city1], []int{city2, L})
		roadInfo[city2] = append(roadInfo[city2], []int{city1, L})
	}
	deepRoad(cur)
	fmt.Println(minRoads, maxTeams+cityTeams[cur])
}

func deepRoad(cur int) {
	if cur == emer {
		if roadLen < minRoadLen {
			minRoadLen = roadLen
			minRoads = 1
			maxTeams = teams
		} else if roadLen == minRoadLen {
			minRoads++
			if teams > maxTeams {
				maxTeams = teams
			}
		}
		return
	}
	for _, v := range roadInfo[cur] {
		if !cityVisited[v[0]] {
			cityVisited[v[0]] = true
		} else {
			continue
		}
		roadLen += v[1]
		teams += cityTeams[v[0]]
		deepRoad(v[0])
		cityVisited[v[0]] = false
		roadLen -= v[1]
		teams -= cityTeams[v[0]]
	}
}
