package main

import (
	"fmt"
	graph "yath/17_graph"
)

func main() {
	var adj graph.AdjucentList
	adj.Init(5)
	adj.AddEdge(0, 1)
	adj.AddEdge(0, 2)
	adj.AddEdge(1, 2)
	adj.AddEdge(1, 3)
	adj.AddEdge(2, 3)
	adj.AddEdge(2, 4)
	adj.AddEdge(3, 4)
	// adj.Print()
	fmt.Println(adj.BFSDisJunction(5))
}
