package graph_test

import (
	"ggyuchive/graph"
	"testing"
)

func checkAdjacencyLengths[T int | graph.WEdge](t *testing.T, name string, g *graph.Graph[T], undirected bool) {
	for ind := range g.Adj {
		expected := 2
		if !undirected {
			expected = 1
		}
		if len(g.Adj[ind]) != expected {
			t.Errorf("[%s] expected len(g.Adj[%d]) = %d, got %d", name, ind, expected, len(g.Adj[ind]))
		}
	}
}

func TestGraphParameterized(t *testing.T) {
	paramTests := []struct {
		name       string
		undirected bool
		weighted   bool
	}{
		{"Unweighted-Undirected", true, false},
		{"Unweighted-Directed", false, false},
		{"Weighted-Undirected", true, true},
		{"Weighted-Directed", false, true},
	}

	for _, tt := range paramTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.weighted {
				g := graph.NewGraph[graph.WEdge](3, tt.undirected)
				g.AddEdge(0, graph.WEdge{1, 3})
				g.AddEdge(1, graph.WEdge{2, 5})
				g.AddEdge(2, graph.WEdge{0, 1})
				checkAdjacencyLengths(t, tt.name, g, tt.undirected)
			} else {
				g := graph.NewGraph[int](3, tt.undirected)
				g.AddEdge(0, 1)
				g.AddEdge(1, 2)
				g.AddEdge(2, 0)
				checkAdjacencyLengths(t, tt.name, g, tt.undirected)
			}
		})
	}
}
