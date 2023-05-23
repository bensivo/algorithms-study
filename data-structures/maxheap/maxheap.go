package maxheap

import "errors"

//     0
//    1     2
//   3 4   5 6
//  7 8 9 10

type MaxHeap struct {
	Items []int
}

func New() *MaxHeap {
	return &MaxHeap{
		Items: []int{},
	}
}

func (mh *MaxHeap) Insert(value int) {
	mh.Items = append(mh.Items, value) // Add the element to the end of the heap

	mh.percolate(len(mh.Items) - 1) // Push the new element up, until the heap property is acheived
}

func (mh *MaxHeap) percolate(index int) {
	if index == 0 {
		return
	}

	pIndex := pIndex(index)

	if mh.Items[pIndex] < mh.Items[index] {
		mh.swap(index, pIndex)
		mh.percolate(pIndex)
	}
}

func (mh *MaxHeap) Pop() (int, error) {
	if len(mh.Items) == 0 {
		return 0, errors.New("heap is empty")
	}

	value := mh.Items[0]

	mh.swap(0, len(mh.Items)-1)           // Copy last element into root
	mh.Items = mh.Items[:len(mh.Items)-2] // Remove last element

	mh.heapify(0) // Pushes the new root down until heap property is acheived again

	return value, nil

}

func (mh *MaxHeap) heapify(index int) {
	li := lIndex(index)
	ri := rIndex(index)
	largest := index

	if li < len(mh.Items) && mh.Items[li] > mh.Items[largest] {
		largest = li
	}
	if ri < len(mh.Items) && mh.Items[ri] > mh.Items[largest] {
		largest = ri
	}

	if largest == index {
		return
	} else {
		mh.swap(index, largest)
		mh.heapify(largest)
	}

}

func (mh *MaxHeap) swap(a int, b int) {
	tmp := mh.Items[a]
	mh.Items[a] = mh.Items[b]
	mh.Items[b] = tmp
}

func pIndex(index int) int {
	return (index - 1) / 2

}
func lIndex(index int) int {
	return 2*index + 1

}
func rIndex(index int) int {
	return 2*index + 2
}
