package main

import (
	"fmt"
	"strconv"
)

func room(N, A, B, C int) []int {
	var res []int
	var r []int
	for i := 0; i < N; i++ {
		res = append(res, i + 1)
	}
	if N == A && A == B && B == C {
		for i := 0; i < N; i++ {
			res[i] = N
		}
		return res
	}
	if N == A && A == B && B != C {
		return r
	}
	if A == 1 && B == 1 {
		return r
	}
	for i := 0; i < A; i++ {
		res[i] = N - A + i + 1
	}
	for i := B - 1; i > 0; i-- {
		res[i] = N - B + (N - i)
	}
	return res
}

func main() {
	var T, N, A, B, C int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&N, &A, &B, &C)
		s := room(N, A, B, C)
		var res string
		if len(s) >= 1 {
			for i := 0; i < len(s); i++ {
				if i == len(s) - 1 {
					res += strconv.Itoa(s[i])
				}
				if i < len(s) - 1 {
					res += strconv.Itoa(s[i]) + " "
				}
			}
			fmt.Printf("Case #%d: %s\n", i+1, res)
		}
		if len(s) < 1 {
			res = "IMPOSSIBLE"
			fmt.Printf("Case #%d: %s\n", i+1, res)
		}
	}
}
