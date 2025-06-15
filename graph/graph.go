package graph

type WEdge struct {
	To   int
	Dist int
}

// Generic type T
type Graph[T int | WEdge] struct {
	N          int
	Undirected bool
	Adj        [][]T
}

// Constructor
func NewGraph[T int | WEdge](n int, undirected bool) *Graph[T] {
	return &Graph[T]{
		N:          n,
		Undirected: undirected,
		Adj:        make([][]T, n),
	}
}

// ensure T is int or struct
// struct T must have field - To int
func (g *Graph[T]) AddEdge(from int, edge T) {
	g.Adj[from] = append(g.Adj[from], edge)
	if g.Undirected {
		switch v := any(edge).(type) {
		case int:
			g.Adj[v] = append(g.Adj[v], any(from).(T))
		case WEdge:
			g.Adj[v.To] = append(g.Adj[v.To], edge)
		default:
			panic("AddEdge: unsupported edge type")
		}
	}
}
