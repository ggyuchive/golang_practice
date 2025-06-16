// solving baekjoon with package (1)
package main

import (
	"fmt"
	"ggyuchive/graph"
	"ggyuchive/input"
)

func solve34009() {
	N := input.NextInt()
	if N%2 == 0 {
		fmt.Print("Alice")
	} else {
		fmt.Print("Bob")
	}
}

func solve34010() {
	N := input.NextInt()
	fmt.Print(N*2-2, N*2-3)
}

func solve34012() {
	N, M := input.NextInt(), input.NextInt()
	g := graph.NewGraph(N, false)
	for i := 0; i < M; i++ {
		a, b := input.NextInt(), input.NextInt()
		g.AddEdge(a, b)
	}
	ans := 0
	for i := 0; i < len(g.Adj); i++ {
		if len(g.Adj[i]) == 1 {
			ans++
		}
	}
	fmt.Print(ans)
}

func solve2644() {
	N := input.NextInt()
	S, E := input.NextInt()-1, input.NextInt()-1
	g := graph.NewGraph(N, true)
	M := input.NextInt()
	for i := 0; i < M; i++ {
		x, y := input.NextInt()-1, input.NextInt()-1
		g.AddEdge(x, y)
	}
	graph.Visited[S], graph.Dist[S] = false, 0
	g.DFS(S)
	if !graph.Visited[E] {
		fmt.Print(-1)
	} else {
		fmt.Print(graph.Dist[E])
	}
}

func main() {
	solve2644()
}
