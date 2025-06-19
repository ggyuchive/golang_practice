package graph

type WEdge struct {
	To   int
	Dist int
}
type WGraph struct {
	N        int
	Directed bool
	Adj      [][]WEdge
	Visited  []bool
	Dist     []int
}

// Weighted Graph Constructor
func NewWGraph(n int, directed bool) *WGraph {
	return &WGraph{
		N:        n,
		Directed: directed,
		Adj:      make([][]WEdge, n),
		Visited:  make([]bool, n),
		Dist:     make([]int, n),
	}
}

func (g *WGraph) AddEdge(from int, edge WEdge) {
	g.Adj[from] = append(g.Adj[from], edge)
	if !g.Directed {
		g.Adj[edge.To] = append(g.Adj[edge.To], WEdge{from, edge.Dist})
	}
}
