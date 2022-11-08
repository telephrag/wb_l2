package develop

import (
	"bufio"
	"log"
	"os"
)

// Sort : TODO
func Sort() {
	file, err := os.Open("input.log") // TODO: pass filename as flag
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

}

func mergeSort(arr []*line) {
	ln := len(arr)

	if ln == 1 {
		return
	}

	pivo := ln / 2

	left, right := arr[0:pivo], arr[pivo:]
	mergeSort(left)
	mergeSort(right)

	merge(arr, pivo)
}

// `arr` is assumed to consist of two non-intersecting sorted arrays
func merge(arr []*line, pivo int) {
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
func bSearch(arr []*line, v *line) int {

	if arr[0].biggerOrEquals(v) {
		return -1 // `v` should be at 0
	}
	if arr[len(arr)-1].smallerOrEquals(v) {
		return len(arr) - 1 // `v` should be appended to `arr`
	}

	i, j := 0, 1
	delta := len(arr) / 2
	for delta > 0 {
		eq := arr[i].equals(v)
		btw := arr[i].smallerOrEquals(v) && v.smallerOrEquals(arr[j]) // `i` and `j` always differ by 1
		big := v.smaller(arr[j])
		sml := v.bigger(arr[i])

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
