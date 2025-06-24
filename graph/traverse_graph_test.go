package graph_test

import (
	"ggyuchive/graph"
	"testing"
)

/*
 *	  4
 *	 /
 *	2 - 1
 *	|
 *	3 - 0	6 - 7 - 5
 */
func createMockGraph() *graph.Graph {
	g := graph.NewGraph(8, false)
	g.AddEdge(4, 2)
	g.AddEdge(2, 1)
	g.AddEdge(2, 3)
	g.AddEdge(3, 0)
	g.AddEdge(6, 7)
	g.AddEdge(7, 5)
	return g
}

func TestDFS(t *testing.T) {
	g := createMockGraph()
	comp := g.DFSAll()
	if comp != 2 {
		t.Errorf("[TestDFS] # of expected component = %d, got %d", 2, comp)
	}
	g.DFS(0)
	expected := []int{0, 3, 2, 1, 3, -1, -1, -1}
	for i := 0; i < len(expected); i++ {
		if g.Dist[i] != expected[i] {
			t.Errorf("Dist[%d]: expected %d, got %d", i, expected[i], g.Dist[i])
		}
	}
}

func TestBFS(t *testing.T) {
	g := createMockGraph()
	comp := g.BFSAll()
	if comp != 2 {
		t.Errorf("[TestDFS] # of expected component = %d, got %d", 2, comp)
	}
	g.BFS(0)
	expected := []int{0, 3, 2, 1, 3, -1, -1, -1}
	for i := 0; i < len(expected); i++ {
		if g.Dist[i] != expected[i] {
			t.Errorf("Dist[%d]: expected %d, got %d", i, expected[i], g.Dist[i])
		}
	}
}
