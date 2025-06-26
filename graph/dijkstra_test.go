package graph_test

import (
	"ggyuchive/graph"
	"testing"
)

/*
 *	     4
 *     /   \
 *   (3)   (2)
 *   /       \
 *  2 - (0) - 1
 *	|
 * (5)
 *	|
 *	3 - (2) - 0    6 - (4) - 7 - (1) - 5
 */
func createMockWGraph() *graph.WGraph {
	g := graph.NewWGraph(8, false)
	g.AddEdge(4, graph.WEdge{2, 3})
	g.AddEdge(2, graph.WEdge{3, 5})
	g.AddEdge(1, graph.WEdge{4, 2})
	g.AddEdge(2, graph.WEdge{1, 0})
	g.AddEdge(3, graph.WEdge{0, 2})
	g.AddEdge(6, graph.WEdge{7, 4})
	g.AddEdge(7, graph.WEdge{5, 1})
	return g
}

func TestDijkstra1(t *testing.T) {
	g := createMockWGraph()
	dist := g.Dijkstra(0)
	expected := []int{0, 7, 7, 2, 9, graph.INF, graph.INF, graph.INF}
	for i := 0; i < len(expected); i++ {
		if dist[i] != expected[i] {
			t.Errorf("Dist[%d]: expected %d, got %d", i, expected[i], g.Dist[i])
		}
	}
}

func TestDijkstra2(t *testing.T) {
	g := createMockWGraph()
	dist := g.Dijkstra(6)
	expected := []int{graph.INF, graph.INF, graph.INF, graph.INF, graph.INF, 5, 0, 4}
	for i := 0; i < len(expected); i++ {
		if dist[i] != expected[i] {
			t.Errorf("Dist[%d]: expected %d, got %d", i, expected[i], g.Dist[i])
		}
	}
}
