package main

import "fmt"

type Graph struct {
	adjMatrix map[int][]int
	visited   map[int]bool
	directed  bool
}

func (g *Graph) initialize(d bool) {
	g.visited = make(map[int]bool)
	g.directed = d
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

func (g *Graph) bfs(n int) {

}

func (g *Graph) dfs(n int) {
	dfsUtil(g, n)
	fmt.Println()
}

func dfsUtil(g *Graph, v int) {
	g.visited[v] = true
	fmt.Printf("%d ", v)

	adjList := g.adjMatrix[v]

	for i := 0; i < len(adjList); i++ {
		if !g.visited[adjList[i]] {
			dfsUtil(g, adjList[i])
		}
	}
}
