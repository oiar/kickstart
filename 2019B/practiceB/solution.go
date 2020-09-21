package main

import (
	"fmt"
	"sort"
)

type stone struct {
	seconds int
	energy  int
	loseEnergy int
}

type stones []stone

func (s stones) Len() int { return len(s) }

func (s stones) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s stones) Less(i, j int) bool {
	return s[i].seconds * s[j].loseEnergy < s[j].seconds * s[i].loseEnergy
}

func energyStones(A []stone) int {
	var res int
	var seconds int
	sort.Sort(stones(A))
	fmt.Println(A)
	for i := 0; i < len(A); i++ {
		seconds += A[i].seconds
	}
	dp := make([]int, seconds + 1)
	for i := 0; i < len(A); i++ {
		for j := seconds; j >= A[i].seconds; j-- {
			e := max(0, A[i].energy - (j - A[i].seconds) * A[i].loseEnergy)
			dp[j] = max(dp[j], dp[j - A[i].seconds] + e)
		}
	}

	for i := 0; i < seconds + 1; i++ {
		res = max(res, dp[i])
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	var T, N int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		var A = make([]stone, N)
		for j := 0; j < N; j++ {
			var s, e, l int
			fmt.Scan(&s, &e, &l)
			A[j].seconds = s
			A[j].energy = e
			A[j].loseEnergy = l
		}
		res := energyStones(A)
		fmt.Printf("Case #%d: %d\n", i+1, res)
	}
}
