package allocation

import (
	"container/heap"
	"testing"

	th "developers-challenge/pkg/testhelpers"
)

func TestMinIntHeap(t *testing.T) {
	testCases := []struct {
		desc   string
		risks  []int
		expOut []int
	}{
		{
			desc:   "Success",
			risks:  []int{2, 77, 30, 20, 10},
			expOut: []int{2, 10, 20, 30, 77},
		},
		{
			desc:   "EmptyRisks_ShouldSucceed",
			risks:  []int{},
			expOut: []int{},
		},
		{
			desc:   "NegativeRisks_ShouldSucceed",
			risks:  []int{2, 77, 30, 20, -10},
			expOut: []int{-10, 2, 20, 30, 77},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mh := initMinHeap(t, tc.risks)
			act := popAllNodes(t, mh)
			th.AssertEqualIntSlices(t, act, tc.expOut)
		})
	}
}

func initMinHeap(t *testing.T, risks []int) *MinHeap {
	t.Helper()

	mh := make(MinHeap, len(risks))
	for i := 0; i < len(risks); i++ {
		mh[i] = DataCenter{risks[i], risks[i]}
	}
	heap.Init(&mh)

	return &mh
}

func popAllNodes(t *testing.T, mh *MinHeap) []int {
	t.Helper()

	result := make([]int, mh.Len())
	i := 0
	for mh.Len() > 0 {
		result[i] = heap.Pop(mh).(DataCenter).actualRisk
		i++
	}

	return result
}
