package main

import (
	"fmt"

	r "developers-challenge/routing"
)

func main() {
	graph := make(map[string][]r.Node)

	graph["A"] = []r.Node{{Id: "B", Latency: 10}, {Id: "C", Latency: 20}}
	graph["B"] = []r.Node{{Id: "D", Latency: 15}}
	graph["C"] = []r.Node{{Id: "D", Latency: 30}}
	graph["D"] = []r.Node{}

	path, latency := r.FindMinimumLatencyPath(graph, []string{}, "A", "D")

	fmt.Printf("path: %v, latency: %v\n", path, latency)
}
