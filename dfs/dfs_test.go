package dfs

import (
	"testing"
)


func TestDepthFirstSearch(t *testing.T) {
  g := Graph{}
  g.AdjList = make(map[int][]int, 0)
  g.AddEdge(Edge{1, 2})
  g.AddEdge(Edge{1, 3})
  g.AddEdge(Edge{2, 1})
  g.AddEdge(Edge{2, 4})
  g.AddEdge(Edge{3, 1})
  g.AddEdge(Edge{3, 4})
  g.AddEdge(Edge{3, 5})
  g.AddEdge(Edge{4, 2})
  g.AddEdge(Edge{4, 3})
  g.AddEdge(Edge{5, 3})
  g.AddEdge(Edge{5, 4})
  g.AddEdge(Edge{6, 4})

  discovered := make(map[int]bool, len(g.AdjList))
  var queue []int
  // Do BFS traversal from all undiscovered nodes
  for i, _ := range g.AdjList {
		if (!discovered[i]) {
      // Mark source vertex as discovered
      discovered[i] = true
      // Push source vertex into the queue
      queue = append(queue, i)
      // Start BFS traversal from vertex i
      dfs(g, i, &queue, discovered)
		}
  }

  if discovered != nil {
		t.Errorf("Expected 5, but it was %v instead.", false)
	}
}