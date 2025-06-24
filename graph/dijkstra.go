package graph

import (
	"container/heap"
	"math"
)

const INF = math.MaxInt64

// Implement min-heap for dijkstra
// Len, Less, Swap, Push, Pop methods are required - it defined as interface
type Item struct {
	node int
	dist int
}
type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func (g *WGraph) Dijkstra(s int) []int {
	g.Dist = make([]int, g.N)
	g.Visited = make([]bool, g.N)

	for i := range g.Dist {
		g.Dist[i] = INF
	}
	g.Dist[s] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{s, 0})

	for pq.Len() > 0 {
		s := heap.Pop(pq).(Item).node
		if g.Visited[s] {
			continue
		}
		g.Visited[s] = true

		for _, edge := range g.Adj[s] {
			e := edge.To
			w := edge.Dist
			if g.Dist[e] > g.Dist[s]+w {
				g.Dist[e] = g.Dist[s] + w
				heap.Push(pq, Item{e, g.Dist[e]})
			}
		}
	}
	return g.Dist
}
