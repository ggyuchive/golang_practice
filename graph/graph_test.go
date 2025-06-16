package graph_test

import (
	"ggyuchive/graph"
	"testing"
)

func checkAdjLength(t *testing.T, name string, g graph.AdjLengthChecker, undirected bool) {
	for ind := range g.Size() {
		expected := 2
		if !undirected {
			expected = 1
		}
		if g.LenAdj(ind) != expected {
			t.Errorf("[%s] expected len(g.Adj[%d]) = %d, got %d", name, ind, expected, g.LenAdj(ind))
		}
	}
}

func TestGraphParameterized(t *testing.T) {
	// Table-Driven Test
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
				g := graph.NewWGraph(3, tt.undirected)
				g.AddEdge(0, graph.WEdge{1, 3})
				g.AddEdge(1, graph.WEdge{2, 5})
				g.AddEdge(2, graph.WEdge{0, 1})
				checkAdjLength(t, tt.name, g, tt.undirected)
			} else {
				g := graph.NewGraph(3, tt.undirected)
				g.AddEdge(0, 1)
				g.AddEdge(1, 2)
				g.AddEdge(2, 0)
				checkAdjLength(t, tt.name, g, tt.undirected)
			}
		})
	}
}
