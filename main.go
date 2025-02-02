package main

import (
	"fmt"
	"math"
)

type Graph struct {
	Vertices []*Vertex
}

type Vertex struct {
	key      int
	adjecent []*Vertex
}

func (g *Graph) AddVertex(key int) {
	if contains(g.Vertices, key) {
		err := fmt.Errorf("vertex %v already exist", key)
		fmt.Println(err.Error())
	} else {
		g.Vertices = append(g.Vertices, &Vertex{key: key})
	}
}

func (g *Graph) getVertex(key int) *Vertex {
	for i, v := range g.Vertices {
		if v.key == key {
			return g.Vertices[i]
		}
	}
	return nil
}

func (g *Graph) AddEdge(k1, k2 int) {
	v1 := g.getVertex(k1)
	v2 := g.getVertex(k2)
	if v1 == nil || v2 == nil {
		err := fmt.Errorf("Vertex not found")
		fmt.Println(err.Error())
	} else if contains(v1.adjecent, k1) || contains(v2.adjecent, k2) {
		err := fmt.Errorf("Edge already exist")
		fmt.Println(err.Error())
	} else {
		v1.adjecent = append(v1.adjecent, v2)
		v2.adjecent = append(v2.adjecent, v1)
	}
}

func (g *Graph) RemoveVertex(k int) {
	v := g.getVertex(k)
	if v == nil {
		err := fmt.Errorf("Vertex not found")
		fmt.Println(err.Error())
	}
	for _, vertex := range v.adjecent {
		for i, adj := range vertex.adjecent {
			if adj.key == k {
				fmt.Println("removing edge", vertex)
				vertex.adjecent = append(vertex.adjecent[:i], vertex.adjecent[i+1:]...)
			}
		}
	}
	for i, vertex := range g.Vertices {
		if vertex.key == k {
			g.Vertices = append(g.Vertices[:i], g.Vertices[i+1:]...)
		}
	}
}

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if v.key == k {
			return true
		}
	}
	return false
}
func (g *Graph) Print() {
	for _, v := range g.Vertices {
		fmt.Printf("\n vertex %v: ", v.key)
		for _, v := range v.adjecent {
			fmt.Printf(" %v ", v.key)
		}
	}
	fmt.Println()

}

func removeEdge(s []*Vertex, k int) []*Vertex {
	for i, v := range s {
		if v.key == k {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
func main() {
	/* 	graph := &Graph{}
	   	for i := 0; i < 5; i++ {
	   		graph.AddVertex(i)
	   	}
	   	graph.AddEdge(1, 2)
	   	graph.AddEdge(1, 4)
	   	graph.AddEdge(1, 3)
	   	graph.AddEdge(1, 4)
	   	graph.Print()
	   	graph.RemoveVertex(2)
	   	graph.Print() */

	gr := NewAdjancyMatrixGraph(5)
	gr.AddEdge(0, 1)
	gr.AddEdge(0, 2)
	gr.AddEdge(1, 3)
	gr.AddEdge(1, 4)

	fmt.Println("Adjacency Matrix:")
	PrintAdjancyMatrix(gr)
	Bfs(gr, 0)
	Dfs(gr, 0, make([]bool, gr.size))

	// Perform BFS from node 0
	//adjancyMatric.AddEdge(1, 2)

	graph := [][]int{
		{0, 10, 15, 20},
		{10, 0, 35, 25},
		{15, 35, 0, 30},
		{20, 25, 30, 0},
	}

	// Number of cities
	cities := len(graph)

	// Array to track visited cities
	visited := make([]bool, cities)

	// Mark the starting city as visited
	visited[0] = true

	// Initialize the answer with a large value
	ans := math.MaxInt32

	// Start TSP from city 0
	tsp(graph, cities, 0, 1, 0, visited, &ans)

	fmt.Printf("The minimum cost of visiting all cities is: %d\n", ans)
}

func tsp(graph [][]int, cities int, pos int, count int, cost int, visited []bool, ans *int) {
	// Base case: If all cities are visited and there is a path back to the starting city
	if count == cities && graph[pos][0] > 0 {
		*ans = min(*ans, cost+graph[pos][0])
		return
	}

	// Traverse all the cities
	for i := 0; i < cities; i++ {
		// If the city is not visited and there is an edge between the current city and city `i`
		if !visited[i] && graph[pos][i] > 0 {
			// Mark the city as visited
			visited[i] = true
			// Recurse with updated parameters
			tsp(graph, cities, i, count+1, cost+graph[pos][i], visited, ans)
			// Backtrack and unmark the city
			visited[i] = false
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
