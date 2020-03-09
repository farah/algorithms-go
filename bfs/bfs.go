package bfs

import "fmt"

type Edge struct {
  Src int
  Dest int
}

type Graph struct {
  AdjList map[int][]int
}


// Iterative method
func bfsIterative(g Graph, v int, discovered map[int]bool) {
  var q []int
  discovered[v] = true
  q = append(q, v)
 
  for len(q) !=0 {
    v, q = q[len(q)-1], q[:len(q)-1]

    for _, u := range g.AdjList[v] {
      if !discovered[u] {
        discovered[u] = true
        q = append(q, u)
      }
    }
    fmt.Println(v)
  }
}

// Recursive method
func bfsRecursive(g Graph, v int, queue *[]int, discovered map[int]bool) {
  if len(*queue) == 0 {
    return
  }
  // Pop vertex from queue
  v, *queue = (*queue)[len(*queue)-1], (*queue)[:len(*queue)-1]

  for _, u := range g.AdjList[v] {
    if !discovered[u] {
      discovered[u] = true
      *queue = append(*queue, u)
    }
  }
  fmt.Println(v)
  bfsRecursive(g, v, queue, discovered);
}


func (g *Graph) AddEdge(edge Edge) { 
   g.AdjList[edge.Src] = append(g.AdjList[edge.Src], edge.Dest)
} 

