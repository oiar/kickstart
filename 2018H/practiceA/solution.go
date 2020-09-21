package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func bigButtons(N int, P []string) float64 {
	n := float64(N)
	res := math.Pow(2, n)
	if len(P) == 1 && len(P[0]) == 1 {
		res -= math.Pow(2, n-1)
		return res
	}
	sort.Sort(ByLength(P))
	var A []string
	for _, v := range P {
		A = append(A, v)
	}
	fmt.Println(P, "P1?")
	for i := 0; i < len(P)-1; i++ {
		for j := i + 1; j < len(P); j++ {
			fmt.Println(P, P[i], P[j], i, j, strings.Contains(P[j], P[i]), strings.Index(P[j], P[i]) == 0, "j, i, con, index")
			if strings.Contains(P[j], P[i]) && strings.Index(P[j], P[i]) == 0 {
				for k, v := range A {
					if P[j] == v {
						A = append(A[:k], A[k+1:]...)
						break
					}
				}
				fmt.Println(A, "A?")
			}
		}
	}
	for _, v := range A {
		r := N - len(v)
		res -= math.Pow(2, float64(r))
	}
	return res
}

func main() {
	var T, N, P int
	var S string
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		fmt.Scan(&P)
		var H []string
		for j := 0; j < P; j++ {
			fmt.Scan(&S)
			H = append(H, S)
		}
		res := bigButtons(N, H)
		fmt.Printf("Case #%d: %d\n", i+1, int(res))
	}
}
