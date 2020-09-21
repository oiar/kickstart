package main

import "fmt"

func alienPiano(K int, A []int) int {
	var res int
	if K < 5 {
		return res
	}
	for i := 0; i < K-4; i++ {
		if compare(A[i], A[i+1], A[i+2], A[i+3], A[i+4]) {
			res++
			i = i + 4
			break
		}
		i = i + 4
	}
	return res
}

func compare(a, b, c, d, e int) bool {
	if e > d && d > c && c > b && b > a {
		return true
	}
	return false
}

func main() {
	var T, K, a int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var A []int
		fmt.Scan(&K)
		for i := 0; i < K; i++ {
			fmt.Scan(&a)
			A = append(A, a)
		}
		s := alienPiano(K, A)
		fmt.Printf("Case #%d: %d\n", i+1, s)
	}
}
