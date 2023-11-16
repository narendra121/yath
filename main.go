package main

import graph "yath/17_graph"

func main() {
	var adj graph.Graph
	adj.Init(5)
	adj.AddEdge(0, 5)
	adj.AddEdge(5, 8)
	adj.AddEdge(5, 8)


	// adj.AddEdge(0, 2)
	// adj.AddEdge(1, 2)
	// adj.AddEdge(1, 3)
	// adj.AddEdge(2, 3)
	// adj.AddEdge(2, 4)
	// adj.AddEdge(3, 4)
	// // adj.Print()
	// fmt.Println(adj.ShortestPath(5,0))

}
