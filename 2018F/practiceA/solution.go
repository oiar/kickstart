package main

import "fmt"

func anagrammatic(L int, A, B string) int {
	var s int
	var x, y []string
	for i := 0; i < L; i++ {
		for j := i + 1; j <= L; j++ {
			a := A[i:j]
			b := B[i:j]
			x = append(x, a)
			y = append(y, b)
		}
	}
	for _, va := range x {
		for _, vb := range y {
			if len(va) == len(vb) {
				res := make([]int, 26)
				h := 0
				// fmt.Println(va, vb)
				for i := 0; i < len(va); i++ {
					// fmt.Println(va[i]-'A', vb[i]-'A')
					res[va[i]-'A']++
					res[vb[i]-'A']--
				}
				for i := 0; i < 26; i++ {
					if res[i] == 0 {
						h++
					}
				}
				// fmt.Println(h)
				if h == 26 {
					s++
					break
				}
			}
		}
	}
	return s
}

func main() {
	var T, L int
	var A, B string
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&L)
		fmt.Scan(&A, &B)
		s := anagrammatic(L, A, B)
		fmt.Printf("Case #%d: %d\n", i+1, s)
	}
}
