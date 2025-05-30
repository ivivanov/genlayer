package allocation

type DataCenter struct {
	baseRisk   int
	actualRisk int
}

type MinHeap []DataCenter

// sort.Interface methods
func (mh MinHeap) Len() int           { return len(mh) }
func (mh MinHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }
func (mh MinHeap) Less(i, j int) bool { return mh[i].actualRisk < mh[j].actualRisk }

// heap.Interface methods
func (mh *MinHeap) Push(x any) { *mh = append(*mh, x.(DataCenter)) }

func (mh *MinHeap) Pop() any {
	deref := *mh
	l := len(deref)
	last := deref[l-1]
	*mh = deref[0 : l-1]

	return last
}
