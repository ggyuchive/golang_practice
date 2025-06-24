package graph

import (
	"container/list"
)

// Get ready for traversal
func (g *Graph) ConstructTraversal(s int) {
	for i := range g.Visited {
		g.Visited[i] = false
		g.Dist[i] = -1
	}
	g.Visited[s], g.Dist[s] = true, 0
}

// DFS using recursion
func (g *Graph) DFS(s int) {
	g.ConstructTraversal(s)
	g.DFS_(s)
}

func (g *Graph) DFS_(s int) {
	for _, e := range g.Adj[s] {
		if !g.Visited[e] {
			g.Visited[e] = true
			g.Dist[e] = g.Dist[s] + 1
			g.DFS_(e)
		}
	}
}

// DFSAll: Traverse all and return # of components
func (g *Graph) DFSAll() int {
	comp := 0
	for s := 0; s < len(g.Adj); s++ {
		if !g.Visited[s] {
			g.DFS(s)
			comp++
		}
	}
	return comp
}

// BFS using queue
func (g *Graph) BFS(s int) {
	var Queue list.List
	g.ConstructTraversal(s)
	Queue.Init().PushBack(s)
	for Queue.Len() > 0 {
		curr := Queue.Remove(Queue.Front()).(int)
		for _, next := range g.Adj[curr] {
			if !g.Visited[next] {
				g.Visited[next] = true
				g.Dist[next] = g.Dist[curr] + 1
				Queue.PushBack(next)
			}
		}
	}
}

// BFSAll: Traverse all and return # of components
func (g *Graph) BFSAll() int {
	comp := 0
	for s := 0; s < len(g.Adj); s++ {
		if !g.Visited[s] {
			g.BFS(s)
			comp++
		}
	}
	return comp
}
