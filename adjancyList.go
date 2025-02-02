package main

import "fmt"

type Graphs struct {
	Vertices map[int][]int
}

func (g *Graphs) AddEdge(from, to int) {
	g.Vertices[from] = append(g.Vertices[from], to)
	g.Vertices[to] = append(g.Vertices[to], from)
}

func (g *Graphs) Print() {
	for key, neighbors := range g.Vertices {
		fmt.Printf("Vertex %d -> ", key)
		for _, v := range neighbors {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}

func NewGraph() *Graphs {
	return &Graphs{Vertices: make(map[int][]int)}
}

func (g *Graphs) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Println(current, ":sss")
		for _, neighbour := range g.Vertices[current] {
			if !visited[neighbour] {
				visited[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}
	fmt.Println()

}

func (g *Graphs) dfsHelper(start int, visited map[int]bool) {
	visited[start] = true
	fmt.Println(start, " ")

	for _, neighbour := range g.Vertices[start] {
		if !visited[neighbour] {
			g.dfsHelper(neighbour, visited)

		}
	}
}

func (g *Graphs) DFS(start int) {
	if _, exist := g.Vertices[start]; !exist {
		fmt.Println("node not exist")
	}

	visited := make(map[int]bool)
	fmt.Print("DFS Traversal: \n")
	g.dfsHelper(start, visited) // Call the recursive function
	fmt.Println()

}
