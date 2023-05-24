package graph_test

import (
	"testing"

	"bensivo.com/leetcode/data-structures/graph"
	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	g := graph.New()

	g.AddNode("A")
	g.AddNode("B")
	g.AddNode("C")
	g.AddNode("D")

	g.AddEdge("A", "B", 1)
	g.AddEdge("A", "C", 2)
	g.AddEdge("C", "D", 1)

	assert.Equal(t, 1, g.Edges["A"]["B"])
	assert.Equal(t, 2, g.Edges["A"]["C"])
}

func TestGraph_DuplicateNode(t *testing.T) {

	g := graph.New()

	g.AddNode("A")
	err := g.AddNode("A")
	assert.EqualError(t, err, "node A already exists")
}

func TestGraph_DuplicateEdge(t *testing.T) {

	g := graph.New()

	g.AddNode("A")
	g.AddNode("B")
	g.AddEdge("A", "B", 1)
	err := g.AddEdge("A", "B", 1)
	assert.EqualError(t, err, "edge A->B already exists")
}
