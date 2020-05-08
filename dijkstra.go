package main

import (
	"math"
)

func dijkstra(graph [][]int, source int) ([]int, []int) {
	var dist []int  // List of distances from source
	var prev []int  // List of previous nodes
	var dSet []bool // List of visited nodes
	for i := 0; i < len(graph); i++ {
		dist = append(dist, math.MaxInt64)
		dSet = append(dSet, false)
		prev = append(prev, 0)
	}
	dist[source] = 0
	for c := 0; c < len(graph); c++ {
		var u int
		u = minDistance(dist, dSet)
		dSet[u] = true
		for v := 0; v < len(graph); v++ {
			if !dSet[v] && graph[u][v] != 0 && dist[u] != math.MaxInt64 && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v]
				prev[v] = u
			}
		}
	}
	return dist, prev
}

func minDistance(dist []int, dset []bool) int {
	var min = math.MaxInt64
	var index = 0
	for v := 0; v < len(dist); v++ {
		if dset[v] == false && dist[v] <= min {
			min = dist[v]
			index = v
		}
	}
	return index
}
