// solving baekjoon with package (1)
package main

import (
	"fmt"
	"ggyuchive/graph"
	"ggyuchive/input"
	"math"
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
	g := graph.NewGraph(N, true)
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
	g := graph.NewGraph(N, false)
	M := input.NextInt()
	for i := 0; i < M; i++ {
		x, y := input.NextInt()-1, input.NextInt()-1
		g.AddEdge(x, y)
	}
	g.DFS(S)
	if !g.Visited[E] {
		fmt.Print(-1)
	} else {
		fmt.Print(g.Dist[E])
	}
}

func solve5014() {
	F, S, G, U, D := input.NextInt(), input.NextInt()-1, input.NextInt()-1, input.NextInt(), input.NextInt()
	g := graph.NewGraph(F, true)
	for a := 0; a < F; a++ {
		if a+U < F {
			g.AddEdge(a, a+U)
		}
		if a-D >= 0 {
			g.AddEdge(a, a-D)
		}
	}
	g.BFS(S)
	if g.Dist[G] == -1 {
		fmt.Print("use the stairs")
	} else {
		fmt.Print(g.Dist[G])
	}
}

func solve13009() {
	N := input.NextInt()
	fmt.Print(int(math.Sqrt(float64(N))))
}

func solve1261() {
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	N, M := input.NextInt(), input.NextInt()
	g := graph.NewWGraph(N*M, true)
	grid := make([][]int, M)
	for i := 0; i < M; i++ {
		str := input.NextStr()
		for j := 0; j < N; j++ {
			grid[i] = append(grid[i], int(str[j]-'0'))
		}
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < 4; k++ {
				ni, nj := i+dx[k], j+dy[k]
				if ni >= 0 && ni < M && nj >= 0 && nj < N {
					g.AddEdge(i+j*M, graph.WEdge{ni + nj*M, grid[ni][nj]})
				}
			}
		}
	}
	ret := g.Dijkstra(0)
	fmt.Print(ret[N*M-1])
}

func main() {
	solve1261()
}
