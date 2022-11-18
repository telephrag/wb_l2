package main

func mergeSort(arr []*Line) {
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
func merge(arr []*Line, pivo int) {

	var sorted []*Line
	if pivo >= len(arr)/2 {
		sorted = arr[:pivo] // part of `arr` left of `pivo`
	} else {
		// swaping arrays
		temp := make([]*Line, pivo)
		copy(temp, arr[:pivo])
		copy(arr, arr[pivo:])
		pivo = len(arr) - pivo
	}

	for i := pivo; i < len(arr); i++ {
		pos := bSearch(sorted, arr[i])

		switch {
		case pos < 0:
			temp := arr[i]          // arr[i] should be leftmost element of `sorted`
			copy(arr[1:], arr[0:i]) // shift right by one
			arr[0] = temp
		case pos == len(sorted)-1: // arr[i] should be rightmost element of `sorted`
			sorted = arr[:i+1]
			continue
		default:
			temp := arr[i]
			copy(arr[pos+2:], arr[pos+1:i])
			arr[pos+1] = temp
		}

		sorted = arr[:i+1]
	}
}

// `arr` is assumed to be pre-sorted
// return value is an index element should be right of
func bSearch(arr []*Line, v *Line) int {

	// Find rightmost element smaller than `v`

	// i := len(arr) / 2
	// delta := i / 2

	// for !btw && !eq {}

	if arr[0].BiggerOrEquals(v) {
		return -1 // `v` should be at 0
	}
	if arr[len(arr)-1].SmallerOrEquals(v) {
		return len(arr) - 1 // `v` should be appended to `arr`
	}

	i, j := 0, 1
	delta := len(arr) / 2
	btw := func() bool {
		return arr[i].SmallerOrEquals(v) && v.SmallerOrEquals(arr[i+1])
	}
	eq := func() bool {
		return arr[i].Equals(v)
	}

	for !eq() && !btw() && j < len(arr) {
		big := v.Smaller(arr[j])
		sml := v.Bigger(arr[i])

		switch {
		case big:
			i, j = i-delta, j-delta
		case sml:
			i, j = i+delta, j+delta
		}
		delta = delta/2 + delta%2
	}

	return i
}
