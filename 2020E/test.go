package main

import (
	"fmt"
)

func Longest(N int, A []int) int {
	var res int
	var B []int
	if N == 2 {
		res = 2
		return res
	}
	for i := 0; i < len(A) - 1; i++ {
		B = append(B, A[i + 1] - A[i])
	}

	var C []int
	var c = 1
	for i := 0; i < len(B) - 1; i++ {
		if B[i] == B[i + 1] {
			c++
		}
		if B[i] != B[i + 1] {
			C = append(C, c)
			c = 1
		}
		if i + 2 == len(B) {
			C = append(C, c)
		}
	}
	if len(C) == 0 {
		res = c + 1
		return res
	}

	var max int
	for i := 0; i < len(C); i++ {
		if C[i] > max {
			max = C[i]
		}
	}
	res = max + 1
	return res
}

func main() {
	var T, N, a int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var A []int
		fmt.Scan(&N)
		for i := 0; i < N; i++ {
			fmt.Scan(&a)
			A = append(A, a)
		}
		s := Longest(N, A)
		fmt.Printf("Case #%d: %d\n", i+1, s)
	}
}
