package graph

func (g *Graph) DFS(s int) {
	if Visited[s] {
		return
	}
	Visited[s] = true
	for _, e := range g.Adj[s] {
		if !Visited[e] {
			Dist[e] = Dist[s] + 1
			g.DFS(e)
		}
	}
}

func (g *Graph) BFS(s int) {
	// TODO
}

// Traverse all components
// Return # of components
func (g *Graph) DFSAll() int {
	ret := 0
	for s := 0; s < len(g.Adj); s++ {
		if !Visited[s] {
			g.DFS(s)
			ret++
		}
	}
	return ret
}

func (g *Graph) BFSAll() int {
	// TODO
	return 0
}
