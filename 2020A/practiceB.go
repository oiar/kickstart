package main

import "fmt"

func plates(A [][]int, P, N, K int) int {
	var res int
	var dp [][]int

	dp = make([][]int, 55)
	for i := 0; i < N + 1; i ++ {
		dp[i] = make([]int, 1555)
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= K; j++ {
			for k := j; k <= P; k++ {
				dp[i][k] = max(dp[i][k], max(dp[i-1][k], dp[i - 1][k - j] + A[i][j]))
			}
		}
	}

	res = dp[N][P]
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	var T, N, K, P, a int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&N, &K, &P)
		var A = make([][]int, 105)
		for j := 1; j <= N; j++ {
			A[j] = make([]int, 105)
			var t int
			for k := 1; k <= K; k++ {
				fmt.Scan(&a)
				A[j][k] = a + t
				t += a
			}
		}
		s := plates(A, P, N, K)
		fmt.Printf("Case #%d: %d\n", i+1, s)
	}
}
