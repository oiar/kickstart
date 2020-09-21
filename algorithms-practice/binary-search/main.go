package main

import (
	"fmt"
	"math"
)

func binarySearch(A []int, s int) int {
	var mid int
	var left, right = 0, len(A)
	for {
		mid = int(math.Floor(float64((left + right) / 2)))
		if A[mid] > s {
			right = mid - 1
		} else if A[mid] < s {
			left = mid + 1
		} else {
			break
		}
		if left > right {
			mid = -1
			break
		}
	}
	return mid
}

func main() {
	var A []int
	var N, a, s int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		A = append(A, a)
	}
	fmt.Scan(&s)
	res := binarySearch(A, s)
	fmt.Println(res)
}
