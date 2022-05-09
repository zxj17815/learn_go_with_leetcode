package main

import "fmt"

var nodes, nonLeafNodes int
var mapTree map[string][]string
var levelMap map[int]int
var level int

func main() {
	levelMap = make(map[int]int)
	_, _ = fmt.Scanf("%d", &nodes)
	_, _ = fmt.Scanf("%d", &nonLeafNodes)
	mapTree = make(map[string][]string)

	for i := 0; i < nonLeafNodes; i++ {
		var children int
		var parent string
		_, _ = fmt.Scanf("%s %d", &parent, &children)
		var childrenList []string
		for j := 0; j < children; j++ {
			var child string
			_, _ = fmt.Scanf("%s", &child)
			childrenList = append(childrenList, child)
		}
		mapTree[parent] = childrenList
	}
	deep("01")
	if nodes == 1 {
		fmt.Printf("%d", 1)
		return
	}
	fmt.Printf("%d", 0)
	for i := 1; i < len(levelMap); i++ {
		fmt.Printf(" %d", levelMap[i])
	}
}

func deep(nodeName string) {
	if _, ok := levelMap[level]; ok == false {
		levelMap[level] = 0
	}
	if _, ok := mapTree[nodeName]; ok {
		level++
		for _, v := range mapTree[nodeName] {
			deep(v)
		}
		level--
	} else {
		levelMap[level]++
		return
	}
}
