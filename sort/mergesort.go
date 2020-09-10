package sort

import (
	"math"
)

// MergeSorter - a type, typically a collection, that satisfies sort.MergeSorter can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
type MergeSorter interface {
	// Len - is the number of elements in the collection.
	Len() int
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
	// Less - reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool

	// For this sorting to work,
	// an additional collection of equal size is needed with the collection being sorted.

	// MemCopy - moves collection elements between left and right  inclusive
	//to the beginning of the additional mem collection
	MemCopy(left, right int)
	// MemLess checks if item at index i is less than item at index j in additional collection
	MemLess(i, j int) bool
	// MemToSlice - moves the element at index memIdx from the mem collection to the main collection at index sliceIdx
	MemToSlice(memIdx, sliceIdx int)
}

// MergeSort - merge sort without recursion
func MergeSort(data MergeSorter) {

	var (
		len = data.Len()
		// numOfPass - maximum number of passing
		numOfPass           = (int)(math.Ceil(math.Log(float64(data.Len())) / math.Log(2)))
		increment           = 2
		left, right, middle int
	)

	for i := 0; i < numOfPass; i++ {
		for j := 0; j < len; j += increment {
			left = j
			if j+increment-1 >= len {
				right = len - 1
			} else {
				right = j + increment - 1
			}
			middle = left + (right-left)/2
			if i == numOfPass-1 {
				middle = left + increment/2 - 1
			}

			merge(data, left, middle, right)
		}
		increment *= 2
	}

}

// merging two parts of the collection
func merge(data MergeSorter, left, middle, right int) {

	if left >= right || middle < left || middle > right {
		return
	}

	if right == left+1 && data.Less(right, left) {
		data.Swap(left, right)
		return
	}

	data.MemCopy(left, right)

	for i, j, k := left, 0, middle-left+1; i <= right; {
		if j > middle-left {
			data.MemToSlice(k, i)
			k++
		} else if k > right-left {
			data.MemToSlice(j, i)
			j++
		} else {
			if data.MemLess(j, k) {
				data.MemToSlice(j, i)
				j++
			} else {
				data.MemToSlice(k, i)
				k++
			}
		}
		i++
	}
}
