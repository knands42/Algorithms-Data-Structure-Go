package dfs

import "fmt"

// Node Represents a node in a graph using adjacency list
type Node[T comparable] struct {
	Val       T
	Neighbors []*Node[T]
	Visited   bool
	Count     int
}

func NewNode[T comparable](val T) *Node[T] {
	return &Node[T]{Val: val}
}

func (n *Node[T]) AddNeighbor(node *Node[T]) {
	n.Neighbors = append(n.Neighbors, node)
}

type Graph[T comparable] struct {
	Nodes []*Node[T]
}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{}
}

func (g *Graph[T]) AddNode(node *Node[T]) *Node[T] {
	g.Nodes = append(g.Nodes, node)
	return node
}

// Dfs O(V + E) time complexity where V is the number of vertices and E is the number of edges
func (g *Graph[T]) Dfs(nodeAt *Node[T]) {
	if nodeAt == nil || nodeAt.Visited {
		return
	}
	nodeAt.Visited = true

	fmt.Println(nodeAt.Val)
	for _, neighbor := range nodeAt.Neighbors {
		g.Dfs(neighbor)
	}
}

// FindComponents O(V + E) time complexity where V is the number of vertices and E is the number of edges
func (g *Graph[T]) FindComponents() int {
	count := 0
	for i := 0; i < len(g.Nodes); i++ {
		if !g.Nodes[i].Visited {
			count++
			g.DfsForConnectedNodes(g.Nodes[i], count)
		}
	}

	return count
}

func (g *Graph[T]) DfsForConnectedNodes(nodeAt *Node[T], count int) {
	nodeAt.Visited = true
	nodeAt.Count = count

	for _, neighbor := range nodeAt.Neighbors {
		if !neighbor.Visited {
			fmt.Println(nodeAt.Val, neighbor.Val, count)
			g.DfsForConnectedNodes(neighbor, count)
		}
	}
}
