package main

import (
	"fmt"
)

type Graph struct {
	adjacency map[int][]int
}

func NewGraph() Graph {
	return Graph{
		adjacency: make(map[int][]int),
	}
}

func (g *Graph) addVertex(vertex int) {
	g.adjacency[vertex] = []int{}
}

func (g *Graph) addEdge(vertex, node int) {
	g.adjacency[vertex] = append(g.adjacency[vertex], node)
}

func (g Graph) createVisited() map[int]bool {
	visited := make(map[int]bool, len(g.adjacency))
	for key := range g.adjacency {
		visited[key] = false
	}

	return visited
}

func (g Graph) bfs(startNode int, visited map[int]bool, res map[int]int) map[int]int{
	var q []int
	q = append(q, startNode)
	for len(q) > 0 {
		var current int
		current, q = q[0], q[1:]
		visited[current] = true
		for _, node := range g.adjacency[current] {
			if !visited[node] {
				res[node] = res[current] + 1
				q = append(q, node)
			}
		}
	}
	return res
}

func planetDistance(g Graph) map[int]int {
	var que []int
	var res = make(map[int]int)
	var mark = make(map[int]bool)
	var degree = make(map[int]int)
	for i := 1; i <= len(g.adjacency); i++ {
		mark[i] = false
		degree[i] = len(g.adjacency[i])
		if degree[i] == 1 {
			que = append(que, i)
		}
	}

	for len(que) > 0 {
		mark[que[0]] = true
		first := que[0]
		que = que[1:]
		next := g.adjacency[first][0]
		degree[next] = degree[next] - 1
		if degree[next] == 1 {
			que = append(que, next)
		}
	}

	visited := g.createVisited()
	for i := 1; i <= len(mark); i++ {
		if !mark[i] {
			res[i] = 0
			visited[i] = true
		}
	}

	for i := 1; i <= len(mark); i++ {
		if !mark[i] {
			g.bfs(i, visited, res)
		}
	}

	return res
}

func main() {
	var T, N, a, b int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		fmt.Scan(&N)
		g := NewGraph()
		for j := 1; j <= N; j++ {
			g.addVertex(j)
		}
		for k := 0; k < N; k++ {
			fmt.Scan(&a, &b)
			g.addEdge(a, b)
			g.addEdge(b, a)
		}
		s := planetDistance(g)

		fmt.Printf("Case #%d:", i + 1)
		for i := 1; i <= len(s); i++ {
			fmt.Printf(" %d", s[i])
		}
		fmt.Printf("\n")
	}
}
