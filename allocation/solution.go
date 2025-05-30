package allocation

import "container/heap"

type DataCenter struct {
	baseRisk   int
	actualRisk int
}

// IncreaseRisk multiplies the DataCenter's actualRisk by its baseRisk.
// This is called when new fragment is stored to the DataCenter.
func (dc *DataCenter) IncreaseRisk() {
	dc.actualRisk = dc.actualRisk * dc.baseRisk
}

type MinHeap []DataCenter

func InitMinHeap(risks []int) *MinHeap {
	mh := make(MinHeap, len(risks))
	for i := 0; i < len(risks); i++ {
		mh[i] = DataCenter{risks[i], risks[i]}
	}
	heap.Init(&mh)

	return &mh
}

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

// Custom methods

// PeekDataCenter returns a pointer to DataCenter with lowest risk without removing it.
func (mh MinHeap) PeekDataCenter() *DataCenter { return &mh[0] }

// StoreFragment stores a new fragment in the DataCenter at the top of the MinHeap,
// increasing its risk and fixing the heap.
// Returns the original risk before the update.
func (mh *MinHeap) StoreFragment() int {
	// retrieve the data center which has minimal acquired risk
	dc := mh.PeekDataCenter()
	// keep the current risk
	originalRisk := dc.actualRisk
	// increase the risk, because new fragment is stored in the data center
	dc.IncreaseRisk()
	// heap element is mutated so the heap needs a fix
	heap.Fix(mh, 0)

	return originalRisk
}

func distributeFragments(risks []int, fragments int) int {
	dataCentersHeap := InitMinHeap(risks)

	// if there is a data center with base risk of 1
	// => we can put all fragments there
	if dataCentersHeap.PeekDataCenter().baseRisk == 1 {
		return 1
	}

	totalRisk := 0
	for i := 0; i < fragments; i++ {
		risk := dataCentersHeap.StoreFragment()
		totalRisk += risk
	}

	return totalRisk
}
