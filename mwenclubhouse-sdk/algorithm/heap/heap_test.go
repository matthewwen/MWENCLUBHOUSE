package heap

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Example struct {
	SortKey int `json:"test"`
}

func (i1 Example) getKey() int {
	return i1.SortKey
}

func (i1 Example) compare(example HeapItem) int {
	return i1.SortKey - example.getKey()
}

func assertSorted(t *testing.T, list []HeapItem) {
	for i := 0; i < len(list)-1; i++ {
		assert.Equal(t, true, list[i].compare(list[i+1]) >= 0)
	}
}

func assertValidHeap(t *testing.T, list []HeapItem, idx int) {
	if idx >= len(list) {
		return
	}
	left, right := 2*idx+1, 2*idx+2
	if left < len(list) {
		assert.Equal(t, true, list[idx].compare(list[left]) <= 0)
		assertValidHeap(t, list, left)
	}
	if right < len(list) {
		assert.Equal(t, true, list[idx].compare(list[right]) <= 0)
		assertValidHeap(t, list, right)
	}

}

func generateList(size int) []HeapItem {
	list := make([]HeapItem, size)
	for i := 0; i < size; i++ {
		list[i] = Example{SortKey: rand.Intn(10000)}
	}
	return list
}

func TestHeapify(t *testing.T) {
	list := []HeapItem{Example{SortKey: 3}, Example{SortKey: 4}, Example{SortKey: -1}, Example{SortKey: 9}}
	Heapify(list)
	assertValidHeap(t, list, 0)
}

func TestHeapSort(t *testing.T) {
	list := []HeapItem{Example{SortKey: 3}, Example{SortKey: 4}, Example{SortKey: -1}, Example{SortKey: 9}}
	HeapSort(list)
	assertSorted(t, list)
}

func TestEmptyList(t *testing.T) {
	var list []HeapItem
	HeapSort(list)
}

func TestHeapFromSize1toSize100(t *testing.T) {
	for i := 1; i < 100; i++ {
		l := generateList(i)
		HeapSort(l)
		assertSorted(t, l)
	}
}

func TestHeapifySize1toSize100(t *testing.T) {
	for i := 1; i < 100; i++ {
		l := generateList(i)
		Heapify(l)
		assertValidHeap(t, l, 0)
	}
}
