package bfs

import (
	"errors"

	"bensivo.com/leetcode/data-structures/graph"
	"bensivo.com/leetcode/data-structures/linkedlist"
	"bensivo.com/leetcode/data-structures/queue"
)

// Given a graph, start, and end nodes, use breadth-first-search to find the path from start to end
func Bfs(g *graph.DirectedGraph, start string, end string) ([]string, error) {

	// Track nodes that have already been visited
	visited := make(map[string]byte)

	// Not strictly a part of BFS, but required to return the path
	// Tracks the parent of each node, so when we get to the end we can build the path backwards
	parent := make(map[string]string)

	// BFS internally uses a queue, at each step we pop from the queue, see if we're at the target, and add all children to the queue
	q := queue.New()

	q.Push(start)
	visited[start] = 0
	for {
		node, err := q.Pop()
		if err == queue.ErrQueueEmpty {
			break
		}

		if node == end { // We've hit the end - now we use the parent map to build our path back to the begining.
			path := linkedlist.New() // We could use a slice here instead of a linkedlist, but then we have to reverse it before returning
			for {
				path.Prepend(node)
				if node == start {
					return path.ToArray(), nil
				}

				node = parent[node]
			}
		}

		// We did not hit the end - go thorugh all reachable nodes from here, and add them to the queue if they haven't already been visisted
		for nextNode := range g.Edges[node] {
			if _, exists := visited[nextNode]; exists {
				continue
			}

			q.Push(nextNode)
			visited[nextNode] = 0
			parent[nextNode] = node
		}
	}

	return nil, errors.New("target not found")
}
