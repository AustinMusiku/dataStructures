package Graphs

import (
	"math"

	"github.com/AustinMusiku/dataStructures/Heap"
)

// Compute the shortest path from a source to all other vertices using Dijkstra's algorithm.
// Returns a map of vertices and their distances from the source.
func (g *Graph) ShortestPath(src *vertex) map[string]float64 {
	distances := make(map[string]float64)
	visited := make(map[string]bool)

	// initialise all distances to a ridiculously high value (+Infinity)
	// except the src which will have 0
	distances[src.key] = 0
	for _, vertex := range g.vertices {
		if vertex.key != src.key {
			distances[vertex.key] = math.Inf(0)
		}
	}

	// Use the pQueue to retrieve the element with the next shortest distance
	priorityQueue := Heap.NewHeap[*vertex]("min")

	// start with the src in the pQueue
	priorityQueue.Insert(src, 0)

	for priorityQueue.Size() > 0 {
		current := priorityQueue.Remove()

		// cycle through all the current vertex's neighbours
		for key, neighbour := range current.Value.neighbours {

			if _, ok := visited[key]; !ok {
				// compare distance already in distances map to
				// distance through the current vertex
				distanceThroughCurrent := distances[current.Value.key] + neighbour.weight
				if distances[key] > distanceThroughCurrent {
					distances[key] = distanceThroughCurrent
					priorityQueue.Insert(neighbour.vertex, int(distanceThroughCurrent))
				}
			}
		}

		visited[current.Value.key] = true
	}

	return distances
}
