package main

import (
	"fmt"
	"testing"
)

func TestMinor(t *testing.T) {
	a := newLine("9188}", 0, " ")
	b := newLine("9065}", 0, " ")
	fmt.Println(a.bigger(b))

	a.line = "9628}"
	fmt.Println(a.bigger(b))
}

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

	// input = []string{
	// 	"1",
	// 	"2",
	// 	"1",
	// 	"1",
	// 	"2",
	// 	"1",
	// 	"6",
	// 	"1",
	// 	"1",
	// 	"2",
	// }
	// lines = make([]*line, len(input))
	// for i, s := range input {
	// 	lines[i] = newLine(s, 1, " ")
	// }
	// mergeSort(lines)
	// for _, l := range lines {
	// 	fmt.Println(l.line)
	// }

	input = []string{
		"{\"thread\": 5, \"timestamp\": 2022-11-08 13:50:17.461054186 +0400 +04 m=+0.006923689, \"result\": 3540}",
		"{\"thread\": 3, \"timestamp\": 2022-11-08 13:50:17.461014326 +0400 +04 m=+0.006883838, \"result\": 2375}",
		"{\"thread\": 4, \"timestamp\": 2022-11-08 13:50:17.460958253 +0400 +04 m=+0.006827765, \"result\": 864}",
		"{\"thread\": 4, \"timestamp\": 2022-11-08 13:50:17.46094216 +0400 +04 m=+0.006811662, \"result\": 7381}",
		"{\"thread\": 4, \"timestamp\": 2022-11-08 13:50:17.460970247 +0400 +04 m=+0.006839749, \"result\": 9663}",
		"{\"thread\": 4, \"timestamp\": 2022-11-08 13:50:17.461078395 +0400 +04 m=+0.006947927, \"result\": 848}",
		"{\"thread\": 3, \"timestamp\": 2022-11-08 13:50:17.461065489 +0400 +04 m=+0.006935001, \"result\": 2464}",
		"{\"thread\": 3, \"timestamp\": 2022-11-08 13:50:17.461053956 +0400 +04 m=+0.006923458, \"result\": 3508}",
		"{\"thread\": 4, \"timestamp\": 2022-11-08 13:50:17.460997261 +0400 +04 m=+0.006866774, \"result\": 9385}",
		"{\"thread\": 3, \"timestamp\": 2022-11-08 13:50:17.461077844 +0400 +04 m=+0.006947346, \"result\": 3316}",
	}
	lines = make([]*line, len(input))
	for i, s := range input {
		lines[i] = newLine(s, 9, " ")
	}
	mergeSort(lines)
	for _, l := range lines {
		fmt.Println(l.getKey())
	}
}

func TestTournamentTree(t *testing.T) { // TODO: Add checks
	input := []*HeapElem{
		{nil, &line{04, 7, "lol kek haha"}},
		{nil, &line{-1, 0, "lmao"}},
		{nil, &line{02, 3, "1 2 3 4 5"}},
		{nil, &line{-1, 0, ""}},
		{nil, &line{5, 11, "some string"}},
		{nil, &line{8, 14, "another string"}},
		{nil, &line{04, 7, "and yet another one"}},
	}

	// expected := []HeapElem{
	// 	{nil, &line{-1, 0, ""}},
	// 	{nil, &line{-1, 0, "lmao"}},
	// 	{nil, &line{04, 7, "lol kek haha"}},
	// 	{nil, &line{5, 11, "some string"}},
	// 	{nil, &line{8, 14, "another string"}},
	// 	{nil, &line{04, 7, "and yet another one"}},
	// }

	buildTournamentTree(input)

}

// func TestBSearch(t *testing.T) { // TODO: add checks
// 	arr := []int{0, 1, 2, 3, 5, 6, 7}
// 	fmt.Println(bSearch(arr, 4))

// 	arr = []int{0, 1, 2, 3}
// 	fmt.Println(bSearch(arr, 4))

// 	arr = []int{5, 6, 7, 8}
// 	fmt.Println(bSearch(arr, 4))
// }
