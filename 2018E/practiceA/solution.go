package main

import (
	"fmt"
)

func main() {
	var T, N, K int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&N, &K)
		var buckets = make(map[int]int)
		var days int
		var max int
		for j := 0; j < N; j++ {
			fmt.Scan(&days)
			buckets[days]++
			if max < days {
				max = days
			}
		}
		var result int

		var drink = 0
		for i := 1; i <= max; i++ {
			drink += K
			if buckets[i] >= drink {
				result += drink
				drink = 0
			} else {
				result += buckets[i]
				drink -= buckets[i]
			}

		}

		fmt.Printf("Case #%d: %d\n", i+1, result)
	}
}
