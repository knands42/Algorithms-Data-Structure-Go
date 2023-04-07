package graph

import "fmt"

type Node[T comparable] struct {
	Val       T
	Neighbors []*Node[T]
	Visited   bool
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
