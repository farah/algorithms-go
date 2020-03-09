package dfs

import "fmt"

type Edge struct {
  Src int
  Dest int
}

type Graph struct {
  AdjList map[int][]int
}

func dfs(g Graph, v int, queue *[]int, discovered map[int]bool) {
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
  
}

func (g *Graph) AddEdge(edge Edge) { 
   g.AdjList[edge.Src] = append(g.AdjList[edge.Src], edge.Dest)
} 

