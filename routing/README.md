# Routing

## Instructions
Run example:
```
go run ./cmd/example/main.go    
```

Run tests:
```
go test -v ./routing
```

## Dependencies

## Explanation
POC based on the assumption the graph is acyclic and directed
I am using Dijkstra's algorithm for finding the shortest distance form vertex 'A' to vertex 'B'. Since the graph is weighted we should always pick the edge with smallest weight - again we can utilize the mechanics of min heap used in the prev task. 

Resources: https://www.geeksforgeeks.org/dijkstras-shortest-path-algorithm-greedy-algo-7/

## Assumptions
I assume the task is to find the fastest path from 'A' to 'B' in directional, cyclical, weighted graph. The type of graph can be determined by the adjacency list in the example:
directional - because A has B as neighbor but B does not have A listed
weighted - because latency can be represented as edge weight
cyclical - because there might be an input where path from 'A' to 'A' (aka self-loop)