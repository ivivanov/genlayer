package routing

import (
	"container/heap"
	"fmt"
	"strings"
)

type neighbor struct {
	id      string
	latency float32 // distance to the neighbor (edge weight)
}

type minHeap []neighbor

// sort.Interface methods
func (mh minHeap) Len() int           { return len(mh) }
func (mh minHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }
func (mh minHeap) Less(i, j int) bool { return mh[i].latency < mh[j].latency }

// heap.Interface methods
func (mh *minHeap) Push(x any) { *mh = append(*mh, x.(neighbor)) }

func (mh *minHeap) Pop() any {
	deref := *mh
	l := len(deref)
	last := deref[l-1]
	*mh = deref[0 : l-1]

	return last
}

type Node struct {
	Id      string
	Latency float32 // distance
}

// FindMinimumLatencyPath computes the path with the minimum total latency between the source and target nodes
// in the given graph. The graph is represented as an adjacency list. The function takes into account a list
// of compressionNodes, which may affect path selection or latency calculations.
// It returns the optimal path as a string and the total latency as a float32.
//
// Parameters:
//   - graph: adjacency list representing the graph
//   - compressionNodes: list of node identifiers that support compression
//   - source: id of the starting node
//   - target: id of the destination node
//
// Returns:
//   - path: formatted path from source to target (e.g. A->B->C)
//   - dist: the total latency of the path
func FindMinimumLatencyPath(graph map[string][]Node, compressionNodes []string, source, target string) (path string, dist float32) {
	// distances map will store the total distance from source router to each router
	distances := make(map[string]float32)

	// distance to source(self) is 0
	distances[source] = 0

	// Traces will keep track every time distance is added/updated,
	// holding the information from where neighbor router is reached.
	// traces[neighbor] = previous_router
	traces := make(map[string]string)

	// setup neighbors min heap
	minHeap := &minHeap{}
	heap.Push(minHeap, neighbor{source, 0})

	// loop until no neighbors are left in the heap
	for minHeap.Len() > 0 {
		// get the "nearest" neighbor
		router := heap.Pop(minHeap).(neighbor)

		// loop through the neighbors and update the shortest distance
		for _, next := range graph[router.id] {
			distanceToNext := distances[router.id] + next.Latency
			_, ok := distances[next.Id]
			// Update if:
			// a) there is no recorded distance to the neighbor
			// b) the current distance is less then the stored one
			if !ok || distanceToNext < distances[next.Id] {
				distances[next.Id] = distanceToNext
				traces[next.Id] = router.id
			}

			// push new neighbor to the heap,
			// so we can keep the loop going until all routers are traversed
			heap.Push(minHeap, neighbor{next.Id, distances[next.Id]})
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
