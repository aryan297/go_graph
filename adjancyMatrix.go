package main

import "fmt"

type AdjancyMatrixGraph struct {
	matrix [][]int
	size   int
}

func AddEdge(g *AdjancyMatrixGraph, i, j int) {
	if i >= 0 && i < g.size && j >= 0 && j < g.size {
		g.matrix[i][j] = 1
		g.matrix[j][i] = 1
	}
}

func NewAdjancyMatrixGraph(size int) *AdjancyMatrixGraph {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}
	return &AdjancyMatrixGraph{matrix: matrix, size: size}
}

func (g *AdjancyMatrixGraph) AddEdge(i, j int) {
	if i >= 0 && i < g.size && j >= 0 && j < g.size {
		g.matrix[i][j] = 1
		g.matrix[j][i] = 1
	}
}

func createAdjancyMatrix(g *AdjancyMatrixGraph) [][]int {
	adjMatrix := make([][]int, g.size)
	for i := range adjMatrix {
		adjMatrix[i] = make([]int, g.size)
	}
	for i := 0; i < g.size; i++ {
		for j := 0; j < g.size; j++ {
			adjMatrix[i][j] = g.matrix[i][j]
		}
	}
	return adjMatrix
}

func PrintAdjancyMatrix(g *AdjancyMatrixGraph) {
	adjMatrix := createAdjancyMatrix(g)
	for _, row := range adjMatrix {
		fmt.Println(row)
	}
}

func Bfs(g *AdjancyMatrixGraph, start int) {
	visited := make([]bool, g.size)
	queue := []int{start}
	visited[start] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Println(node, "node")
		for i := 0; i < g.size; i++ {
			if g.matrix[node][i] == 1 && !visited[i] {
				queue = append(queue, i)
				visited[i] = true
			}
		}
	}
	fmt.Println("BFS done", visited)
}

func Dfs(g *AdjancyMatrixGraph, start int, visited []bool) {
	visited[start] = true
	fmt.Println(start, "node")
	for i := 0; i < g.size; i++ {
		if g.matrix[start][i] == 1 && !visited[i] {
			Dfs(g, i, visited)
		}
	}
}
