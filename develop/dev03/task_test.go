package develop

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	Sort()
}

func TestMergeSort(t *testing.T) { // TODO: add checks
	arr := []string{
		"lol kek haha",
		"lmao",
		"",
		"1 2 3 4 5",
		"some string",
		"another string",
		"and yet another one",
	}

	lines := make([]*line, len(arr))
	for i, s := range arr {
		lines[i] = newLine(s, 1, " ")
	}

	mergeSort(lines)

	for _, l := range lines {
		fmt.Printf("%+v\n", l)
	}
}

// func TestBSearch(t *testing.T) { // TODO: add checks
// 	arr := []int{0, 1, 2, 3, 5, 6, 7}
// 	fmt.Println(bSearch(arr, 4))

// 	arr = []int{0, 1, 2, 3}
// 	fmt.Println(bSearch(arr, 4))

// 	arr = []int{5, 6, 7, 8}
// 	fmt.Println(bSearch(arr, 4))
// }
