package routing

import (
	"container/heap"
	"fmt"
)

type Neighbor struct {
	id      string
	latency int // distance to the neighbor (edge weight)
}

type MinHeap []Neighbor

// sort.Interface methods
func (mh MinHeap) Len() int           { return len(mh) }
func (mh MinHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }
func (mh MinHeap) Less(i, j int) bool { return mh[i].latency < mh[j].latency }

// heap.Interface methods
func (mh *MinHeap) Push(x any) { *mh = append(*mh, x.(Neighbor)) }

func (mh *MinHeap) Pop() any {
	deref := *mh
	l := len(deref)
	last := deref[l-1]
	*mh = deref[0 : l-1]

	return last
}

type Node struct {
	id      string
	latency int // distance
}

// Dijkstra
func findMinimumLatencyPath(graph map[string][]Node, compressionNodes []string, source, target string) string {
	// distances map will store the total distance from source router to each router
	distances := make(map[string]int)
	// distance to source(self) is 0
	distances[source] = 0

	// setup neighbors min heap
	minHeap := &MinHeap{}
	heap.Push(minHeap, Neighbor{source, 0})

	// loop until no neighbors are left in the heap
	for minHeap.Len() > 0 {
		// get the "nearest" neighbor
		router := heap.Pop(minHeap).(Neighbor)

		// loop through the neighbors and update the shortest distance
		for _, neighbor := range graph[router.id] {
			distanceToNeighbor := distances[router.id] + neighbor.latency
			_, ok := distances[neighbor.id]
			// Update if:
			// a) there is no recorded distance to the neighbor
			// b) the current distance is less then the stored one
			if !ok || distanceToNeighbor < distances[neighbor.id] {
				distances[neighbor.id] = distanceToNeighbor
			}

			// push new neighbor to the heap, 
			// so we can keep the loop going until all routers are traversed
			heap.Push(minHeap, Neighbor{neighbor.id, distances[neighbor.id]})
		}
	}

	fmt.Println(distances)
	return ""
}
