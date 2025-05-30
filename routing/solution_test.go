package routing

import (
	"testing"

	th "developers-challenge/pkg/testhelpers"
)

func TestFindMinimumLatencyPath(t *testing.T) {
	testCases := []struct {
		desc             string
		graph            map[string][]Node
		compressionNodes []string
		source           string
		target           string
		expPath          string
		expLatency       float32 // distance
	}{
		{
			desc: "Success_With_3Routers_WithoutCompression",
			graph: func() map[string][]Node {
				graph := make(map[string][]Node)

				graph["A"] = []Node{{"B", 10}, {"C", 20}}
				graph["B"] = []Node{{"D", 15}}
				graph["C"] = []Node{{"D", 30}}
				graph["D"] = []Node{}

				return graph

			}(),
			source:     "A",
			target:     "D",
			expPath:    "A->B->D",
			expLatency: 25,
		},
		// {
		// 	desc: "Success_With_3Routers_WithCompression",
		// 	graph: func() map[string][]Node {
		// 		graph := make(map[string][]Node)

		// 		graph["A"] = []Node{{"B", 10}, {"C", 20}}
		// 		graph["B"] = []Node{{"D", 15}}
		// 		graph["C"] = []Node{{"D", 30}}
		// 		graph["D"] = []Node{}

		// 		return graph

		// 	}(),
		// 	source:           "A",
		// 	target:           "D",
		// 	compressionNodes: []string{"B", "C"},
		// 	expPath:          "A->B->D",
		// 	expLatency:       17.5,
		// },
		{
			desc: "Success_With_3Routers_WithoutCompression",
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
			expPath:          "A->C->D",
			expLatency:       10,
		},
		{
			desc: "Success_With_4Routers_WithoutCompression",
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
			expPath:          "A->B->F->D",
			expLatency:       6,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actPath, actLatency := FindMinimumLatencyPath(tc.graph, tc.compressionNodes, tc.source, tc.target)
			th.AssertEqualStrings(t, actPath, tc.expPath)
			th.AssertEqualFloats(t, actLatency, tc.expLatency)
		})
	}
}

func TestTraceBack(t *testing.T) {
	testCases := []struct {
		desc   string
		traces map[string]string
		target string
		expOut []string
	}{
		{
			desc: "Success_With_3Routers",
			traces: func() map[string]string {
				t := make(map[string]string)
				t["B"] = "A"
				t["C"] = "A"
				t["D"] = "C"
				t["E"] = "B"
				return t
			}(),
			target: "D",
			expOut: []string{"D", "C", "A"},
		},
		{
			desc: "Success_With_3Routers",
			traces: func() map[string]string {
				t := make(map[string]string)
				t["B"] = "A"
				t["C"] = "A"
				t["D"] = "C"
				t["E"] = "B"
				return t
			}(),
			target: "E",
			expOut: []string{"E", "B", "A"},
		},
		{
			desc: "Success_With_2Routers",
			traces: func() map[string]string {
				t := make(map[string]string)
				t["B"] = "A"
				t["C"] = "A"
				t["D"] = "C"
				t["E"] = "B"
				return t
			}(),
			target: "B",
			expOut: []string{"B", "A"},
		},
		{
			desc: "WhenTarget_IsSource_ShouldReturnEmpty_NoLoopEdge",
			traces: func() map[string]string {
				t := make(map[string]string)
				t["B"] = "A"
				t["C"] = "A"
				t["D"] = "C"
				t["E"] = "B"
				return t
			}(),
			target: "A",
			expOut: []string{},
		},
		// {
		// 	desc: "WhenTarget_IsSource_ShouldReturnSelf_WithLoopEdge",
		// 	traces: func() map[string]string {
		// 		t := make(map[string]string)
		// 		t["A"] = "A"
		// 		t["B"] = "A"
		// 		t["C"] = "A"
		// 		t["D"] = "C"
		// 		t["E"] = "B"
		// 		return t
		// 	}(),
		// 	target: "A",
		// 	expOut: []string{"A"},
		// },
		{
			desc: "WhenNoPath_ToTarget_ShouldReturnEmpty",
			traces: func() map[string]string {
				t := make(map[string]string)
				t["B"] = "A"
				t["C"] = "A"
				t["D"] = "C"
				t["E"] = "B"
				return t
			}(),
			target: "F",
			expOut: []string{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			act := traceBack(tc.traces, tc.target)
			th.AssertEqualStringSlices(t, act, tc.expOut)
		})
	}
}

func TestPrettyPrintPath(t *testing.T) {
	testCases := []struct {
		desc         string
		backwardPath []string
		expOut       string
	}{
		{
			desc:         "Success",
			backwardPath: []string{"C", "B", "A"},
			expOut:       "A->B->C",
		},
		{
			desc:         "Success_When_OneElement",
			backwardPath: []string{"A"},
			expOut:       "A",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			act := prettyPrintPath(tc.backwardPath)
			th.AssertEqualStrings(t, act, tc.expOut)
		})
	}
}
