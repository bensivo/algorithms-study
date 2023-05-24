package graph

import (
	"fmt"
)

//	 There are generally 3 ways to represent a graph in memory:
//		- Adjacency Matrix - NxN matrix, where each row/column is a node, and values represent edge weights - useful for small graphs
//		- Adjacency List - Store an array of node structs, and each node contains an list of other nodes (representing edges to those nodes) - useful for sparse graphs
//		- Edge list - Just store lists to edges. Node existence is assumed by its existence in the edge list
type DirectedGraph struct {
	// This representation is similar to adjacency lists, except it uses maps instead of arrays (for constant time access)
	//
	// Interpretation: g.Edges["A"]["B"] = 1, there is an edge with weight 1 from node "A" to node "B"
	Edges map[string]map[string]int
}

func New() *DirectedGraph {
	return &DirectedGraph{
		Edges: make(map[string]map[string]int), // Usage: g.edges["source node"]["dest node"] = weight of edge
	}
}

func (g *DirectedGraph) AddNode(name string) error {
	_, exists := g.Edges[name]
	if exists {
		return fmt.Errorf("node %s already exists", name)
	}

	g.Edges[name] = make(map[string]int)

	return nil
}

func (g *DirectedGraph) AddEdge(sourceName string, destName string, weight int) error {
	_, exists := g.Edges[sourceName]
	if !exists {
		return fmt.Errorf("node %s does not exist", sourceName)
	}

	_, exists = g.Edges[destName]
	if !exists {
		return fmt.Errorf("node %s does not exist", destName)
	}

	_, exists = g.Edges[sourceName][destName]
	if !exists {
		return fmt.Errorf("edge %s->%s already exists", sourceName, destName)
	}

	edgesOfSource := g.Edges[sourceName]
	edgesOfSource[destName] = weight

	return nil
}
