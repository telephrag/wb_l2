package develop

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Line struct {
	Key   string
	Words []string
}

func NewLine(str string, key int, sep string) *Line {
	l := &Line{}
	l.Words = strings.Split(str, sep)
	if len(l.Words) > key {
		l.Key = l.Words[key]
	}
	return l
}

func Sort() {
	file, err := os.Open("input.txt") // TODO: pass filename as flag
	if err != nil {
		log.Fatal(err)
	}

	lines := make([]*Line, 0)
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)
	read := 0
	for s.Scan() {
		t := s.Text()
		read += len(t)
		lines = append(lines, NewLine(t, 0, " "))
	}

	fmt.Println(lines)
}

func MergeSort() {
	arr0 := []int{23, 17, 38, 10, 190, 123, 5}
	arr1 := []int{627, 27, 13, 87, 11, 12, 1}

	mergeSort(arr0)
	mergeSort(arr1)
	fmt.Println(arr0)

	pivo := len(arr0)
	arr0 = append(arr0, arr1...)
	mergeSorted(arr0, pivo)
	fmt.Println(arr0)
}

func mergeSort(arr []int) {
	ln := len(arr)

	if ln == 1 {
		return
	}

	pivo := ln / 2

	left, right := arr[0:pivo], arr[pivo:]
	mergeSort(left)
	mergeSort(right)

	mergeSorted(arr, pivo)
}

// `arr` is assumed to consist of two non-intersecting sorted arrays
func mergeSorted(arr []int, pivo int) {
	sorted := arr[0:pivo]              // part of `arr` left of `pivo`
	for i := pivo; i < len(arr); i++ { // TODO: Move to seperate function merge and reuse.
		pos := bSearch(sorted, arr[i])

		switch {
		case pos < 0:
			temp := arr[i]                   // arr[i] should be leftmost element of `sorted`
			copy(arr[1:], arr[:len(sorted)]) // shift right by one
			arr[0] = temp
		case pos == len(sorted)-1: // arr[i] should be rightmost element of `sorted`
			continue
		default:
			temp := arr[i]
			copy(arr[pos+2:], arr[pos+1:len(sorted)])
			arr[pos+1] = temp
		}

		sorted = arr[:i+1]
	}
}

// `arr` is assumed to be pre-sorted
// return value is an index element should be right of
func bSearch(arr []int, v int) int {

	if arr[0] >= v {
		return -1 // `v` should be at 0
	}
	if arr[len(arr)-1] <= v {
		return len(arr) - 1 // `v` should be appended to `arr`
	}

	i, j := 0, 1
	delta := len(arr) / 2
	for delta > 0 {
		eq := arr[i] == v
		btw := (arr[i] <= v && v <= arr[j]) // `i` and `j` always differ by 1
		big := v < arr[j]
		sml := v > arr[i]

		switch {
		case eq || btw:
			return i
		case big:
			i, j = i-delta, j-delta
		case sml:
			i, j = i+delta, j+delta
		}
		delta /= 2
	}

	return i
}
