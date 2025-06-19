package graph_test

import (
	"ggyuchive/graph"
	"testing"
)

func TestGraphParameterized(t *testing.T) {
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
			g := graph.NewGraph(3, tt.directed)
			g.AddEdge(0, 1)
			g.AddEdge(1, 2)
			g.AddEdge(2, 0)
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
