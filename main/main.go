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

func solve34013() {
	N := input.NextInt()
	g := graph.NewGraph[graph.WEdge](N*N, true)
	g.AddEdge(1, graph.WEdge{To: 1, Dist: 1})
}

func main() {
	solve34013()
}
