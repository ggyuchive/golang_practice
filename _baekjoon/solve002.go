// solving baekjoon with package (2)
package main

import (
	"fmt"
	"ggyuchive/graph"
	"ggyuchive/input"
)

func solve4485() {
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	for i := 1; ; i++ {
		N := input.NextInt()
		if N == 0 {
			break
		}
		g := graph.NewWGraph(N*N, true)
		grid := make([][]int, N)
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				grid[i] = append(grid[i], input.NextInt())
			}
		}
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				for k := 0; k < 4; k++ {
					ni, nj := i+dx[k], j+dy[k]
					if ni >= 0 && ni < N && nj >= 0 && nj < N {
						g.AddEdge(i+j*N, graph.WEdge{ni + nj*N, grid[ni][nj]})
					}
				}
			}
		}
		ret := g.Dijkstra(0)
		fmt.Printf("Problem %d: %d\n", i, ret[N*N-1]+grid[0][0])
	}
}

func main() {
	solve4485()
}
