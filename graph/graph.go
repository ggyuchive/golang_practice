package graph

var (
	Visited []bool
	Dist    []int
)

type AdjLengthChecker interface {
	LenAdj(i int) int
	Size() int
}
type Graph struct {
	N          int
	Undirected bool
	Adj        [][]int
}

// Graph Constructor
func NewGraph(n int, undirected bool) *Graph {
	Visited = make([]bool, n)
	Dist = make([]int, n)
	return &Graph{
		N:          n,
		Undirected: undirected,
		Adj:        make([][]int, n),
	}
}

func (g *Graph) AddEdge(from int, to int) {
	g.Adj[from] = append(g.Adj[from], to)
	if g.Undirected {
		g.Adj[to] = append(g.Adj[to], from)
	}
}

func (g *Graph) LenAdj(i int) int {
	return len(g.Adj[i])
}
func (g *Graph) Size() int {
	return g.N
}

type WEdge struct {
	To   int
	Dist int
}
type WGraph struct {
	N          int
	Undirected bool
	Adj        [][]WEdge
}

// WGraph Constructor
func NewWGraph(n int, undirected bool) *WGraph {
	Visited = make([]bool, n)
	Dist = make([]int, n)
	return &WGraph{
		N:          n,
		Undirected: undirected,
		Adj:        make([][]WEdge, n),
	}
}

func (g *WGraph) AddEdge(from int, edge WEdge) {
	g.Adj[from] = append(g.Adj[from], edge)
	if g.Undirected {
		g.Adj[edge.To] = append(g.Adj[edge.To], WEdge{from, edge.Dist})
	}
}

func (g *WGraph) LenAdj(i int) int {
	return len(g.Adj[i])
}
func (g *WGraph) Size() int {
	return g.N
}
