package develop

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	Sort()
}

func TestMergeSort(t *testing.T) {
	MergeSort()
}

func TestBSearch(t *testing.T) {
	arr := []int{0, 1, 2, 3, 5, 6, 7}
	fmt.Println(bSearch(arr, 4))

	arr = []int{0, 1, 2, 3}
	fmt.Println(bSearch(arr, 4))

	arr = []int{5, 6, 7, 8}
	fmt.Println(bSearch(arr, 4))
}
