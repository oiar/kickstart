package main

import "fmt"

func maxRopeLength(A []int, K int) int {
	var res int

	return res
}

func main() {
	var A []int
	var a, N, K int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		A = append(A, a)
	}
	fmt.Scan(&K)
	res := maxRopeLength(A, K)
	fmt.Println(res)
}
