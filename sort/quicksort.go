package sort

import (
	"sync"
)

// QuickSorter - a type, typically a collection, that satisfies sort.QuickSorter can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
type QuickSorter interface {
	// Len - is the number of elements in the collection.
	Len() int
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
	// Less - reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
}

// QuickSort - recursive function call wrapper
func QuickSort(data QuickSorter) {
	quickSort(data, 0, data.Len()-1, nil)
}

// quickSort multithreaded hoare quicksort implementation
func quickSort(data QuickSorter, left, right int, wg *sync.WaitGroup) {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if left < right {
		p := partition(data, left, right)

		var wait sync.WaitGroup
		wait.Add(2)
		go quickSort(data, left, p, &wait)
		go quickSort(data, p+1, right, &wait)
		wait.Wait()
	}
}

// hoare tiling
func partition(data QuickSorter, left, right int) int {
	pivot, i, j := left, left, right
	for i <= j {
		for ; data.Less(i, pivot); i++ {
		}
		for ; j > left && !data.Less(j, pivot); j-- {
		}

		if i >= j {
			break
		}

		if pivot == i {
			pivot = j
		} else if pivot == j {
			pivot = i
		}
		data.Swap(i, j)
		i++
		j--
	}

	return j
}
