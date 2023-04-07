package graph

import "testing"

func Test_Dfs(t *testing.T) {
	g := NewGraph[int]()

	n0 := NewNode(0)
	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)

	g.AddNode(n0)
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)

	n0.AddNeighbor(n1)
	n0.AddNeighbor(n2)
	n1.AddNeighbor(n2)
	n2.AddNeighbor(n0)
	n2.AddNeighbor(n3)
	n3.AddNeighbor(n3)

	g.Dfs(n1)
}
