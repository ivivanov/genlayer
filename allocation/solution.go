package allocation

import "container/heap"

type dataCenter struct {
	baseRisk   int
	actualRisk int
}

// IncreaseRisk multiplies the DataCenter's actualRisk by its baseRisk.
// This is called when new fragment is stored to the DataCenter.
func (dc *dataCenter) IncreaseRisk() {
	dc.actualRisk = dc.actualRisk * dc.baseRisk
}

type minHeap []dataCenter

func initMinHeap(risks []int) *minHeap {
	mh := make(minHeap, len(risks))
	for i := 0; i < len(risks); i++ {
		mh[i] = dataCenter{risks[i], risks[i]}
	}
	heap.Init(&mh)

	return &mh
}

// sort.Interface methods
func (mh minHeap) Len() int           { return len(mh) }
func (mh minHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }
func (mh minHeap) Less(i, j int) bool { return mh[i].actualRisk < mh[j].actualRisk }

// heap.Interface methods
func (mh *minHeap) Push(x any) { *mh = append(*mh, x.(dataCenter)) }

func (mh *minHeap) Pop() any {
	deref := *mh
	l := len(deref)
	last := deref[l-1]
	*mh = deref[0 : l-1]

	return last
}

// Custom methods

// peekDataCenter returns a pointer to DataCenter with lowest risk without removing it.
func (mh minHeap) peekDataCenter() *dataCenter { return &mh[0] }

// storeFragment stores a new fragment in the DataCenter at the top of the MinHeap,
// increasing its risk and fixing the heap.
// Returns the original risk before the update.
func (mh *minHeap) storeFragment() int {
	// retrieve the data center which has minimal acquired risk
	dc := mh.peekDataCenter()
	// keep the current risk
	originalRisk := dc.actualRisk
	// increase the risk, because new fragment is stored in the data center
	dc.IncreaseRisk()
	// heap element is mutated so the heap needs a fix
	heap.Fix(mh, 0)

	return originalRisk
}

// DistributeFragments distributes the fragments among data centers
// represented by their risk values. Returns the minimized maximum risk
// after all fragments have been distributed.
//
// Parameters:
//   - risks: a slice of integers representing the initial risk values of each data center.
//   - fragments: the number of fragments to distribute.
//
// Returns:
//   - The minimized maximum risk value after distribution.
func DistributeFragments(risks []int, fragments int) int {
	dataCentersHeap := initMinHeap(risks)

	// if there is a data center with base risk of 1
	// => we can put all fragments there
	if dataCentersHeap.peekDataCenter().baseRisk == 1 {
		return 1
	}

	totalRisk := 0
	for i := 0; i < fragments; i++ {
		risk := dataCentersHeap.storeFragment()
		totalRisk += risk
	}

	return totalRisk
}
