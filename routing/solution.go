package routing

import (
	"container/heap"
	"fmt"
	"strings"
)

type Neighbor struct {
	id      string
	latency float32 // distance to the neighbor (edge weight)
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
	latency float32 // distance
}

// Returns the path
func findMinimumLatencyPath(graph map[string][]Node, compressionNodes []string, source, target string) (path string, dist float32) {
	// distances map will store the total distance from source router to each router
	distances := make(map[string]float32)

	// distance to source(self) is 0
	distances[source] = 0

	// Traces will keep track every time distance is added/updated,
	// holding the information from where neighbor router is reached.
	// traces[neighbor] = previous_router
	traces := make(map[string]string)

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
				traces[neighbor.id] = router.id
			}

			// push new neighbor to the heap,
			// so we can keep the loop going until all routers are traversed
			heap.Push(minHeap, Neighbor{neighbor.id, distances[neighbor.id]})
		}
	}

	path = prettyPrintPath(traceBack(traces, target))
	dist = distances[target]
	return path, dist
}

// traceBack will rebuild the path from "source" to "target"
// using the traces map.
// Returns the path in backward order from "target" to "source"
func traceBack(traces map[string]string, target string) []string {
	// No path to target
	if _, ok := traces[target]; !ok {
		return []string{}
	}

	// restore backward path
	backwardPath := []string{target}

	for {
		prev, ok := traces[target]
		if ok {
			backwardPath = append(backwardPath, prev)
			target = prev
		} else {
			break
		}
	}

	return backwardPath
}

// prettyPrintPath using the backwardPath array will
// reverse the path in the format: A->B->C
// Returns formatted path in proper order
func prettyPrintPath(backwardPath []string) string {
	var sb strings.Builder
	for i := len(backwardPath) - 1; i >= 0; i-- {
		if i == 0 {
			sb.WriteString(backwardPath[i])
		} else {
			sb.WriteString(fmt.Sprintf("%v->", backwardPath[i]))
		}
	}

	return sb.String()
}
