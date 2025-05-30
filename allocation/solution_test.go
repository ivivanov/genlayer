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
			mh := InitMinHeap(tc.risks)
			act := popAllNodes(t, mh)
			th.AssertEqualIntSlices(t, act, tc.expOut)
		})
	}
}

func TestDistributeFragments(t *testing.T) {
	testCases := []struct {
		desc      string
		risks     []int
		fragments int
		expRisk   int
	}{
		{
			desc:      "Success#1",
			risks:     []int{20, 10, 2, 15},
			fragments: 3,
			expRisk:   14,
		},
		{
			desc:      "Success#2",
			risks:     []int{10, 20, 30},
			fragments: 5,
			expRisk:   560,
		},
		{
			desc:      "AllFragments_InOneDataCenter_ShouldSucceed",
			risks:     []int{10, 20, 30, 1},
			fragments: 500,
			expRisk:   1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actRisk := distributeFragments(tc.risks, tc.fragments)
			th.AssertEqualInts(t, actRisk, tc.expRisk)
		})
	}
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
