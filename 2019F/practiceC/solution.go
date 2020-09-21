package main

import "fmt"

type Node struct {
	value int
	children []*Node
}

func newNode(value int, children []*Node) *Node {
	return &Node{
		value: value,
		children: children,
	}
}

func main() {
	var T, V, b, l, r int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&V)
		node := make([]*Node, 10)
		for j := 0; j < V; j++ {
			fmt.Scan(&b)
			child := make([]*Node, V)
			node[j + 1] = newNode(b, child)
		}
		for k := 0 ; k < V - 1; k++ {
			fmt.Scan(&l, &r)
			node[l].children = append(node[l].children, node[r])
		}
		fmt.Println(node[1], node[1].value,  "-------node-------")
	}
}