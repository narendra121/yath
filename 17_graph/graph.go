package graph

import (
	"fmt"
	"math"
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

type Graph struct {
	adj [][]int
}

func (g *Graph) Init(size int) {
	g.adj = make([][]int, size)
}

func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

func (g *Graph) Print() {
	for _, edges := range g.adj {
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

func (g *Graph) BFSDisJunction(v int) int {
	count := 0
	visited := make([]bool, v+1)
	for i := 0; i < v; i++ {
		if !visited[i] {
			g.BreadthFirstTree(i, visited)
			count++
		}
	}
	return count
}

func (g *Graph) BreadthFirstTree(s int, visited []bool) {
	var q queue.QueueArrImpl
	visited[s] = true
	q.Init(0, 10)
	q.EnQueue(s)
	for !q.IsEmpty() {
		u := q.Arr[q.GetFrontIdx().(int)]
		fmt.Println(u.(int))
		q.DeQueue()
		for _, val := range g.adj[u.(int)] {
			if !visited[val] {
				visited[val] = true
				q.EnQueue(val)
			}
		}
	}
}

func (g *Graph) DFS(vertex, start int) int {
	visited, count := make([]bool, vertex), 0
	for i := 0; i < vertex; i++ {
		if !visited[i] {
			g.DFSRecursion(i, visited)
			count++
		}
	}
	return count
}

func (g *Graph) DFSRecursion(source int, visited []bool) {
	visited[source] = true
	fmt.Println(source)
	for _, val := range g.adj[source] {
		if !visited[val] {
			g.DFSRecursion(val, visited)
		}
	}
}

func (g *Graph) ShortestPath(vertices, source int) []int {
	distance := make([]int, vertices)
	distance[source] = 0
	var q queue.QueueArrImpl
	q.Init(0, vertices)
	visited := make([]bool, vertices)
	q.EnQueue(source)
	visited[source] = true
	for !q.IsEmpty() {
		u := q.Arr[q.GetFrontIdx().(int)]
		q.DeQueue()
		for _, val := range g.adj[u.(int)] {
			if !visited[val] {
				distance[val] = distance[u.(int)] + 1
				visited[val] = true
				q.EnQueue(val)
			}
		}
	}
	return distance
}

/*
Prims algorithm (find minimum weight of edges to connect all vertices)
maintain two two sets
one set will have all connected vertices
second will have all non connected vertices
*/

/*
|0 5 8 0   |
|5 0 10 0  |
|8 10 0 20 |
|0 15 20 0 |

op:- 28
*/

func (g *Graph) PrimsAlgo(v int) int {
	key, res := make([]int, v), 0
	for i, _ := range key {
		key[i] = math.MaxInt
	}
	key[0] = 0
	mSet := make([]bool, v)

	for count := 0; count < v; count++ {
		u := -1
		for i := 0; i < v; i++ {
			if !mSet[i] && (u == -1 || key[i] < key[u]) {
				u = i
			}
		}

		mSet[u] = true
		res = res + key[u]

		for k := 0; k < v; k++ {
			if g.adj[u][k] != 0 && !mSet[k] {
				key[k] = int(math.Min(float64(key[k]), float64(g.adj[u][k])))
			}
		}
	}
	return res

}
