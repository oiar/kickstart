package main

import "fmt"

type graph struct {
	numNodes int
	edges    [][]int
}

func newGraph(n int) *graph {
	return &graph{
		numNodes: n,
		edges:    make([][]int, n),
	}
}

// undirected graph
func (g *graph) addEdge(adj [][]int) {
	for i := 0; i < g.numNodes; i++ {
		for j := 0; j < g.numNodes; j++ {
			g.edges[i][j] = -1
		}
	}
	for i := 0; i <= g.numNodes; i++ {
		l := adj[i][0]
		for j := 0; j <= g.numNodes; j++ {
			g.edges[i] = append(g.edges[i], 0)
			g.edges[0] = append(g.edges[0], i)
			if l > adj[i][j] {
				g.edges[i+1] = append(g.edges[i+1], j+1)
				g.edges[j+1] = append(g.edges[j+1], i+1)
				l = adj[i][j]
			}
		}
	}
}

func (g *graph) adjacencyMatrix(from, to, len int) [][]int {
	var adj = make([][]int, g.numNodes)
	for i := 0; i < g.numNodes; i++ {
		for j := 0; j < g.numNodes; j++ {
			adj[i][j] = -1
		}
	}
	adj[from][to] = len
	adj[to][from] = len
	return adj
}

func specializingVillages() int {
	var res int
	return res
}

func main() {
	var T int
	var V, E int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var A, B, L int
		fmt.Scan(&V)
		g := newGraph(V)
		fmt.Scan(&E)
		var adj [][]int
		for j := 0; j < E; j++ {
			fmt.Scan(&A, &B, &L)
			adj = g.adjacencyMatrix(A, B, L)
		}
		g.addEdge(adj)
		// s := anagrammatic(L, A, B)
		// fmt.Printf("Case #%d: %d\n", i+1, s)
	}
}
