package heap

type HeapItem interface {
	getKey() int
	compare(item HeapItem) int
}

func swap(items []HeapItem, idx1, idx2 int) {
	i1, i2 := items[idx1], items[idx2]
	items[idx2], items[idx1] = i1, i2
}

func AddToHeap(items []HeapItem, len *int, idx int) {
	swap(items, *len, idx)
	idx = *len
	*len += 1

	for idx-1 >= 0 {
		var parentIdx int = (idx - 1) / 2
		if items[parentIdx].compare(items[idx]) > 0 {
			swap(items, parentIdx, idx)
			idx = parentIdx
		} else {
			break
		}
	}
}

func PopFromHeap(items []HeapItem, heapLen *int) {
	swap(items, *heapLen-1, 0)
	*heapLen -= 1
	for idx := 0; true; {
		root1Idx, root2Idx := 2*idx+1, 2*idx+2
		swapIdx := -1
		if root1Idx < *heapLen || root2Idx < *heapLen {
			if root1Idx >= *heapLen {
				swapIdx = root2Idx
			} else if root2Idx >= *heapLen {
				swapIdx = root1Idx
			} else {
				if items[root1Idx].compare(items[root2Idx]) < 0 {
					swapIdx = root1Idx
				} else {
					swapIdx = root2Idx
				}
			}
			if items[idx].compare(items[swapIdx]) > 0 {
				swap(items, idx, swapIdx)
				idx = swapIdx
			} else {
				return
			}
		} else {
			return
		}
	}

}

func Heapify(items []HeapItem) {
	heapSize := 0
	for i := 0; i < len(items); i++ {
		AddToHeap(items, &heapSize, i)
	}
}

func HeapSort(items []HeapItem) {
	Heapify(items)
	heapSize := len(items)
	for i := 0; i < len(items); i++ {
		PopFromHeap(items, &heapSize)
	}

}
