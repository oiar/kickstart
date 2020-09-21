// package main

// import (
// 	"fmt"
// 	"math"
// )

// func noNine(F, L int) int {
// 	var res int
// 	for i := F; i <= L; i++ {
// 		if i%9 == 0 {
// 			continue
// 		}
// 		s := i
// 		t := 0
// 		for s > 0 {
// 			// fmt.Println(t, s, "t, s")
// 			t = s % 10
// 			s = s / 10
// 			// fmt.Println(t, s, "t, s")
// 			if t == 9 {
// 				res--
// 				break
// 			}
// 		}
// 		res++
// 		// fmt.Println(res, "res")
// 	}
// 	return res
// }

package main

import (
	"fmt"
)

func noNine(t int) int {
	var A []int
	var digit [18]int
	var res int
	var n = t
	for i := 1; i < 18; i++ {
		digit[0] = 1
		digit[i] = digit[i-1] * 9
	}
	for n > 0 {
		i := n % 10
		n = n / 10
		A = append(A, i)
	}
	for i := len(A) - 1; i > 0; i-- {
		res += A[i] * digit[i-1] * 8
		if A[i] == 9 {
			return res
		}
	}
	t -= A[0]
	for i := 0; i <= A[0] && i != 9; i++ {
		if ((t + i) % 9) > 0 {
			res++
		}
	}
	return res
}

func main() {
	var T, F, L int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&F)
		fmt.Scan(&L)
		s := noNine(L) - noNine(F) + 1
		fmt.Printf("Case #%d: %d\n", i+1, s)
	}
}
