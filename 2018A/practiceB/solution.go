package main

import (
	"fmt"
)

func luckyDip(A []float64, K int) float64 {
	var res float64
	for i := 0; i < len(A); i++ {
		res += A[i]
	}
	res /= float64(len(A))
	// fmt.Println(res)
	if K == 0 {
		return float64(res)
	}
	for i := 0; i < K; i++ {
		var e float64
		count := 0
		for j := 0; j < len(A); j++ {
			if A[j] > res {
				e += A[j]
				count++
			}
		}
		res = (e + float64(len(A)-count)*res) / float64(len(A))
		// fmt.Println(res, "?")
	}
	return float64(res)
}

func main() {
	var T, N, K int
	var n int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&N, &K)
		var A []float64
		for j := 0; j < N; j++ {
			fmt.Scan(&n)
			A = append(A, float64(n))
		}
		s := luckyDip(A, K)
		fmt.Printf("Case #%d: %f\n", i+1, s)
	}
}
