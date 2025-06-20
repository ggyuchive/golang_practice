package graph_test

import (
	"ggyuchive/graph"
	"testing"
)

func TestWGraphParameterized(t *testing.T) {
	// Table-Driven Test
	paramTests := []struct {
		name     string
		directed bool
	}{
		{"Directed", true},
		{"Undirected", false},
	}

	for _, tt := range paramTests {
		t.Run(tt.name, func(t *testing.T) {
			g := graph.NewWGraph(3, tt.directed)
			g.AddEdge(0, graph.WEdge{1, 2})
			g.AddEdge(1, graph.WEdge{2, 3})
			g.AddEdge(2, graph.WEdge{0, 1})
			for ind := range len(g.Adj) {
				expected := 2
				if tt.directed {
					expected = 1
				}
				if len(g.Adj[ind]) != expected {
					t.Errorf("[%s] expected len(g.Adj[%d]) = %d, got %d", tt.name, ind, expected, len(g.Adj[ind]))
				}
			}
		})
	}
}
