package main

import (
	"fmt"
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

func main() {
	solve34009()
}
