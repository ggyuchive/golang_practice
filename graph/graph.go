package graph

type Graph struct {
	N        int
	Directed bool
	Adj      [][]int
	Visited  []bool
	Dist     []int
}

// Graph Constructor
func NewGraph(n int, directed bool) *Graph {
	return &Graph{
		N:        n,
		Directed: directed,
		Adj:      make([][]int, n),
		Visited:  make([]bool, n),
		Dist:     make([]int, n),
	}
}

func (g *Graph) AddEdge(from int, to int) {
	g.Adj[from] = append(g.Adj[from], to)
	if !g.Directed {
		g.Adj[to] = append(g.Adj[to], from)
	}
}
