package graph

import (
	"fmt"
	queue "yath/12_queue"
)

/*
G(v,e)
v--> vertices  (each node ) (v1,v2,v3)
e--> edges (pairs) ((v1,v2),(v2,v3))

Directed graph --> it will have direction from vertex to other using edges

Undirected -->it is random

Degree of vertex --> no of edges is going through it (lines) (for undirected)

Indegree and out degree (directed graph)

walk --> from one vertex to any other which is connected with edges..

cyclic --> begin and ends with same vertex

Weighted graph --> weight of edges defined by distance between two vertices

Adjacent matrix --- vertex * vertex matrix  and fill the ones based on index directions

adj.Init(5)
adj.AddEdge(0, 1)
adj.AddEdge(0, 2)
adj.AddEdge(1, 2)
adj.AddEdge(1, 3)
adj.AddEdge(2, 3)
adj.AddEdge(2, 4)
adj.AddEdge(3, 4)
adj.BreadthFirstTree(5, 0)

*/

type AdjucentList struct {
	adj [][]int
}

func (a *AdjucentList) Init(size int) {
	a.adj = make([][]int, size)
}

func (a *AdjucentList) AddEdge(u, v int) {
	a.adj[u] = append(a.adj[u], v)
	a.adj[v] = append(a.adj[v], u)
}

func (a *AdjucentList) Print() {
	for _, edges := range a.adj {
		for _, edge := range edges {
			fmt.Print(edge, "\t")
		}
		fmt.Println()
	}
}

/*
BFS
v is size of vertex
s is source vertex
returning number of islands or connected components
*/

func (a *AdjucentList) BFSDisJunction(v int) int {
	count := 0
	visited := make([]bool, v+1)
	for i := 0; i < v; i++ {
		if !visited[i] {
			a.BreadthFirstTree(i, visited)
			count++
		}
	}
	return count
}
func (a *AdjucentList) BreadthFirstTree(s int, visited []bool) {
	var q queue.QueueArrImpl
	visited[s] = true
	q.Init(0, 10)
	q.EnQueue(s)
	for !q.IsEmpty() {
		u := q.Arr[q.GetFrontIdx().(int)]
		fmt.Println(u.(int))
		q.DeQueue()
		for _, val := range a.adj[u.(int)] {
			if !visited[val] {
				visited[val] = true
				q.EnQueue(val)
			}
		}
	}
}
