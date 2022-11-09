package main

import (
	"testing"
)

func TestMergeSort(t *testing.T) { // TODO: add checks
	input := []string{
		"lol kek haha",
		"lmao",
		"",
		"1 2 3 4 5",
		"some string",
		"another string",
		"and yet another one",
	}

	lines := make([]*line, len(input))
	for i, s := range input {
		lines[i] = newLine(s, 1, " ")
	}

	mergeSort(lines)

	expected := []string{
		"",
		"lmao",
		"1 2 3 4 5",
		"lol kek haha",
		"some string",
		"another string",
		"and yet another one",
	}

	for i, l := range lines {
		if l.line != expected[i] {
			t.Errorf("expected \"%s\" at position %d; received \"%s\"",
				expected[i],
				i,
				l.line,
			)
		}
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
