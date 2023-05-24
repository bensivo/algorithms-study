package bfs_test

import (
	"testing"

	"bensivo.com/leetcode/algorithms/bfs"
	"bensivo.com/leetcode/data-structures/graph"
	"github.com/stretchr/testify/assert"
)

func TestBfs_Graph1(t *testing.T) {
	g := graph.New()

	g.AddNode("A")
	g.AddNode("B")
	g.AddNode("C")
	g.AddNode("D")
	g.AddNode("E")
	g.AddNode("F")

	g.AddEdge("A", "B", 0)
	g.AddEdge("B", "C", 0)
	g.AddEdge("C", "A", 0)
	g.AddEdge("C", "D", 0)
	g.AddEdge("D", "E", 0)
	g.AddEdge("E", "A", 0)
	g.AddEdge("D", "F", 0)

	path, _ := bfs.Bfs(g, "A", "F")

	assert.Equal(t, []string{"A", "B", "C", "D", "F"}, path)
}

func TestBfs_Graph1_NotFound(t *testing.T) {
	g := graph.New()

	g.AddNode("A")
	g.AddNode("B")
	g.AddNode("C")
	g.AddNode("D")
	g.AddNode("E")
	g.AddNode("F")

	g.AddEdge("A", "B", 0)
	g.AddEdge("B", "C", 0)
	g.AddEdge("C", "A", 0)
	g.AddEdge("C", "D", 0)
	g.AddEdge("D", "E", 0)
	g.AddEdge("E", "A", 0)
	g.AddEdge("D", "F", 0)

	_, err := bfs.Bfs(g, "A", "G")
	assert.EqualError(t, err, "target not found")
}
