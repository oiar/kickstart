package main

import (
	"fmt"
	"math"
)

func evenDigit(N int) int {
	var res, increase, decrease, sum int
	var digit []int
	var s = N
	for s > 0 {
		t := s % 10
		s = s / 10
		digit = append(digit, t)
	}

	for i := len(digit) - 1; i >= 0; i-- {
		if digit[i] == 9 {
			de := digit[i] - 1
			for j := i - 1; j >= 0; j-- {
				decrease += 8 * int(math.Pow10(j))
			}
			decrease += int(math.Pow10(i) * float64(de)) + sum
			return N - decrease
		}
		if digit[i] != 9 && digit[i] % 2 == 1 {
			in := digit[i] + 1
			de := digit[i] - 1
			increase = int(math.Pow10(i) * float64(in)) + sum
			
			for j := i - 1; j >= 0; j-- {
				decrease += 8 * int(math.Pow10(j))
			}
			decrease += int(math.Pow10(i) * float64(de)) + sum

			break
		}
		sum += int(math.Pow10(i) * float64(digit[i]))
	}
	a := increase - N
	b := N - decrease
	if a > 0 && b > 0 {
		if a >= b {
			res = b
		} else {
			res = a
		}
	}
	return res
}

func main() {
	var T, N int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		s := evenDigit(N)
		fmt.Printf("Case #%d: %d\n", i+1, s)
	}
}