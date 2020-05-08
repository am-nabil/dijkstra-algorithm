package main

import (
	"fmt"
	"testing"
)

func TestDijkstra(t *testing.T) {
	var graph = [][]int{
		{0, 5, 3, 0, 0, 0, 0},
		{0, 0, 2, 0, 3, 0, 1},
		{0, 0, 0, 7, 7, 0, 0},
		{2, 0, 0, 0, 0, 6, 0},
		{0, 0, 0, 2, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0}}
	const source = 0

	dist, prev := dijkstra(graph, source)

	for i := 0; i < len(dist); i++ {
		fmt.Printf("Distance from %d to %d is : %d \n", source, i, dist[i])
	}
	for i := 0; i < len(dist); i++ {
		fmt.Printf("Previous node of %d from %d is : %d \n", i, source, prev[i])
	}

	if dist[0] != 0 || dist[1] != 5 || dist[2] != 3 || dist[3] != 9 || dist[4] != 7 || dist[5] != 8 || dist[6] != 6 {
		t.Errorf("Distances incorrect")
	}
	if prev[0] != 0 || prev[1] != 0 || prev[2] != 0 || prev[3] != 4 || prev[4] != 6 || prev[5] != 4 || prev[6] != 1 {
		t.Errorf("Previous nodes incorrect")
	}
}
