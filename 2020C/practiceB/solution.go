package main

import (
	"fmt"
)

type Node struct {
	name     string
	children []*Node
}

func newNode(name string) *Node {
	return &Node{
		name:     name,
		children: []*Node{},
	}
}

func (n *Node) addChild(o *Node) {
	n.children = append(n.children, o)
}

func topSort(g []*Node) []*Node {
	visited := make(map[string]bool)
	order := make([]*Node, 0, len(g))
	for _, n := range g {
		if visited[n.name] == false {
			visited[n.name] = true
			visitedSubTree(n, visited, &order)
		}
	}
	return reverse(order)
}

func visitedSubTree(n *Node, visited map[string]bool, order *[]*Node) {
	*order = append(*order, n)
	for _, v := range n.children {
		if visited[v.name] == false {
			visited[v.name] = true
			visitedSubTree(v, visited, order)
		}
	}
}

func reverse(order []*Node) []*Node {
	for i, j := 0, len(order); i < j; i, j = i+1, j-1 {
		order[i], order[j] = order[j], order[i]
	}
	return order
}

func main() {
	var T, R, C int
	var S string
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&R)
		for j := 0; j < R; j++ {
			fmt.Scan(&C)
			for k := 0; k < C; k++ {
				fmt.Scan(&S)
			}
		}
		// r, e := graph.topSort(S)
		// if e == -1 {
		// 	fmt.Printf("Case #%d: %d\n", i+1, e)
		// }
		// fmt.Printf("Case #%d: %d\n", i+1, r)
	}
}
