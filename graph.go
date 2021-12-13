package main

import "fmt"

type Graph struct {
	adjMatrix map[int][]int
	visited   map[int]int
	directed  bool
}

func (g *Graph) initialize() {
	g.visited = make(map[int]int)
}

func (g *Graph) addEdge(a int, b int) {
	if g.adjMatrix == nil {
		g.adjMatrix = make(map[int][]int)
	}

	if g.adjMatrix[a] == nil {
		g.adjMatrix[a] = []int{}
	}

	g.adjMatrix[a] = append(g.adjMatrix[a], b)

	if !g.directed {
		if g.adjMatrix[b] == nil {
			g.adjMatrix[b] = []int{}
		}
		g.adjMatrix[b] = append(g.adjMatrix[b], a)
	}
}

func (g *Graph) dfs(n int) {
	if g.visited[n] == 0 {
		fmt.Printf("%d->", n)
		g.visited[n] = 1

		for _, j := range g.adjMatrix[n] {
			if g.visited[j] == 0 {
				g.dfs(j)
			}
		}
	}
}
