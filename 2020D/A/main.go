package main

import "fmt"

func main() {
	var T, N int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {

		// var V = make([]int, 200000)
		fmt.Scan(&N)
		var max int
		var previous int
		var current int
		var isIncrease bool
		var result int
		if N == 1 {
			fmt.Scan(&current)
			fmt.Printf("Case #%d: %d\n", i+1, 1)
			continue
		}
		for j := 0; j < N; j++ {
			fmt.Scan(&current)

			if current >= max {

				if current > max {
					if current > previous {
						if !isIncrease {
							result++
							isIncrease = true
						}

					}
				}

				if current == max {
					if isIncrease && current == previous {
						result--
						isIncrease = false
					}
				}
				max = current
			}

			if current < previous {
				isIncrease = false
			}

			previous = current

		}
		fmt.Printf("Case #%d: %d\n", i+1, result)
	}
}
