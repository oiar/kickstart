package main

import (
	"fmt"
	"math"
)

type Stack struct {
	nodes []int
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n int) {
	s.nodes = append(s.nodes, n)
	s.count++
}

func (s *Stack) Pop() int {
	var t int
	if s.count != 0 {
		s.count--
		t = s.nodes[s.count]
		s.nodes = s.nodes[:s.count]
	}
	return t
}

func direction(w, h, num int, r rune) (int, int) {
	if r == 'E' {
		w += num
	}
	if r == 'S' {
		h += num
	}
	if r == 'N' {
		h -= num
	}
	if r == 'W' {
		w -= num
	}
	return w, h
}

func checkout(w, h int) (int, int) {
	s := math.Pow10(9) - 1
	if float64(w+1) > math.Pow10(9) {
		w = int(s) % int(math.Pow10(9))
	}
	if float64(h+1) > math.Pow10(9) {
		h = int(s) % int(math.Pow10(9))
	}
	if w+1 < 1 {
		w = int(s) + w + 1
	}
	if h+1 < 1 {
		h = int(s) + h + 1
	}
	return w + 1, h + 1
}

func robotPath(K string) (int, int) {
	var w, h int
	num := 1
	s := NewStack()
	s.Push(1)
	for _, v := range K {
		if v >= '0' && v <= '9' {
			num = num * int(v-'0')
			s.Push(int(v - '0'))
			// fmt.Println(s, "s")
		}
		if v == ')' {
			t := s.Pop()
			// fmt.Println(s, "what is s")
			num /= t
		}
		// fmt.Println(w, h, v, num)
		w, h = direction(w, h, num, v)
	}
	// fmt.Println(w, h, "w, h")
	w, h = checkout(w, h)
	return w, h
}

func main() {
	var T int
	var K string
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&K)
		w, h := robotPath(K)
		fmt.Printf("Case #%d: %d %d\n", i+1, w, h)
	}
}
