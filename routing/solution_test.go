package routing

import "testing"

func TestFindMinimumLatencyPath(t *testing.T) {
	testCases := []struct {
		desc             string
		graph            map[string][]Node
		compressionNodes []string
		source           string
		target           string
		expOut           string
		expRisk          int
	}{
		{
			desc: "S",
			graph: func() map[string][]Node {
				graph := make(map[string][]Node)

				graph["A"] = []Node{{"B", 10}, {"C", 20}}
				graph["B"] = []Node{{"D", 15}}
				graph["C"] = []Node{{"D", 30}}
				graph["D"] = []Node{}

				return graph

			}(),
			// compressionNodes: []string{"A", "B"},
			source: "A",
			target: "D",
			expOut: "A->B->D",
		},
		{
			desc: "",
			graph: func() map[string][]Node {
				graph := make(map[string][]Node)

				graph["A"] = []Node{{"B", 4}, {"C", 8}}
				graph["B"] = []Node{{"E", 6}}
				graph["C"] = []Node{{"D", 2}}
				graph["D"] = []Node{{"E", 10}}

				return graph

			}(),
			compressionNodes: []string{"A", "B"},
			source:           "A",
			target:           "D",
			expOut:           "A->C->D",
		},
		{
			desc: "",
			graph: func() map[string][]Node {
				graph := make(map[string][]Node)

				graph["A"] = []Node{{"B", 4}, {"C", 8}}
				graph["B"] = []Node{{"E", 6}, {"F", 1}}
				graph["C"] = []Node{{"D", 2}}
				graph["D"] = []Node{{"E", 10}}
				graph["E"] = []Node{}
				graph["F"] = []Node{{"D", 1}}

				return graph

			}(),
			compressionNodes: []string{"A", "B"},
			source:           "A",
			target:           "D",
			expOut:           "A->B->F->D",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			path := findMinimumLatencyPath(tc.graph, tc.compressionNodes, tc.source, tc.target)
			t.Log(path)
		})
	}
}
