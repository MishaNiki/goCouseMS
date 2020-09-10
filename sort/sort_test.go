package sort

import (
	"testing"
)

type testSort []int

var (
	tmp []int
)

func (ts testSort) Len() int              { return len(ts) }
func (ts testSort) Swap(i, j int)         { ts[i], ts[j] = ts[j], ts[i] }
func (ts testSort) Less(i, j int) bool    { return ts[i] < ts[j] }
func (ts testSort) MemLess(i, j int) bool { return tmp[i] < tmp[j] }
func (ts testSort) MemCopy(left, right int) {
	for i := 0; i < right-left+1; i++ {
		tmp[i] = ts[left+i]
	}
}
func (ts testSort) MemToSlice(memIdx, sliceIdx int) { ts[sliceIdx] = tmp[memIdx] }

func testArrays() []testSort {
	return []testSort{
		{129, 323, 212, 434, 534834, 2332, 0, -1, 233, 43, 454, 5332, 44, 22, 34, 5454, 54, 323, 11, 0, 12, 1111, 4444},
		{100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 0},
		{0, 1, 2, 3, 4, 5},
		{0, 0},
		{0, 1},
	}
}

func TestMergeSort(t *testing.T) {
	ta := testArrays()
	for i := 0; i < len(ta); i++ {
		tmp = make([]int, len(ta[i]))
		MergeSort(ta[i])
		for j := 1; j < len(ta[i]); j++ {
			if ta[i][j-1] > ta[i][j] {
				t.Errorf("Merge sort no sorted, testArray: %v", i)
			}
		}
	}
}

func TestQuickSort(t *testing.T) {
	ta := testArrays()
	for i := 0; i < len(ta); i++ {
		QuickSort(ta[i])
		for j := 1; j < len(ta[i]); j++ {
			if ta[i][j-1] > ta[i][j] {
				t.Errorf("QuickSort no sorted, testArray: %v", i)
			}
		}
	}
}

func TestHeapSort(t *testing.T) {

	ta := testArrays()
	for i := 0; i < len(ta); i++ {
		HeapSort(ta[i])
		for j := 1; j < len(ta[i]); j++ {
			if ta[i][j-1] > ta[i][j] {
				t.Errorf("HeapSort no sorted, testArray: %v", i)
			}
		}
	}
}
