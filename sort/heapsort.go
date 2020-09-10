package sort

// HeapSorter - a type, typically a collection, that satisfies sort.HeapSorter can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
type HeapSorter interface {
	// Len - is the number of elements in the collection.
	Len() int
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
	// Less - reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
}

// HeapSort ...
func HeapSort(data HeapSorter) {
	len := data.Len()
	for j := 0; j < len; j++ {
		for i := len/2 - 1 - j/2; i > -1; i-- {
			if 2*i+2 <= len-1-j {
				if !data.Less(2*i+1, 2*i+2) {
					if data.Less(i, 2*i+1) {
						data.Swap(i, 2*i+1)
					}
				} else if data.Less(i, 2*i+2) {
					data.Swap(i, 2*i+2)
				}
			} else if 2*i+1 <= len-1-j && data.Less(i, 2*i+1) {
				data.Swap(i, 2*i+1)
			}
		}
		data.Swap(0, len-1-j)
	}
}
