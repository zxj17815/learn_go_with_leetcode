package main

import "fmt"

func main() {
	sum := 1.0
	result := [...]string{"W", "T", "L"}
	for i := 0; i < 3; i++ {
		var maxOdd float64
		var r string
		for _, v := range result {
			var tempOdd float64
			_, _ = fmt.Scanf("%f", &tempOdd)
			if tempOdd > maxOdd {
				maxOdd = tempOdd
				r = v
			}
		}
		fmt.Printf("%v ", r)
		sum = sum * maxOdd
	}
	fmt.Printf("%.2f", (sum*0.65-1)*2)
}
