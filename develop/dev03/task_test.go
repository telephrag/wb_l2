package main

import (
	"testing"
)

// func TestMinor(t *testing.T) {
// 	a := newLine("9188}", 0, -1, " ")
// 	b := newLine("9065}", 0, -1, " ")
// 	fmt.Println(a.Bigger(b))

// 	a.line = "9628}"
// 	fmt.Println(a.Bigger(b))
// }

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

	lines := make([]*Line, len(input))
	for i, s := range input {
		lines[i] = newLine(s, 1, i, " ")
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
	lines = make([]*Line, len(input))
	for i, s := range input {
		lines[i] = newLine(s, 9, i, " ")
	}
	mergeSort(lines)
	// for _, l := range lines {
	// 	fmt.Println(l.GetKey())
	// }
}

// func TestBSearch(t *testing.T) { // TODO: add checks
// 	arr := []int{0, 1, 2, 3, 5, 6, 7}
// 	fmt.Println(bSearch(arr, 4))

// 	arr = []int{0, 1, 2, 3}
// 	fmt.Println(bSearch(arr, 4))

// 	arr = []int{5, 6, 7, 8}
// 	fmt.Println(bSearch(arr, 4))
// }

func TestTournamentTree(t *testing.T) { // TODO: Add checks
	input := []string{
		"{\"thread\": 6, \"timestamp\": 2022-11-08 13:50:17.460267568 +0400 +04 m=+0.006137081, \"result\": 7698}", // these two
		"{\"thread\": 2, \"timestamp\": 2022-11-08 13:50:17.460682015 +0400 +04 m=+0.006551527, \"result\": 2440}", //
		"{\"thread\": 6, \"timestamp\": 2022-11-08 13:50:17.460784652 +0400 +04 m=+0.006654164, \"result\": 7308}", // then this one
		"{\"thread\": 5, \"timestamp\": 2022-11-08 13:50:17.460492983 +0400 +04 m=+0.006362486, \"result\": 7289}",
		"{\"thread\": 2, \"timestamp\": 2022-11-08 13:50:17.460304443 +0400 +04 m=+0.006173945, \"result\": 7174}",
		"{\"thread\": 2, \"timestamp\": 2022-11-08 13:50:17.460608236 +0400 +04 m=+0.006477738, \"result\": 7086}",
		"{\"thread\": 5, \"timestamp\": 2022-11-08 13:50:17.460803701 +0400 +04 m=+0.006673213, \"result\": 2572}",
		"{\"thread\": 1, \"timestamp\": 2022-11-08 13:50:17.46064465 +0400 +04 m=+0.006514282, \"result\": 7082}",
	}

	sources := make([]*LineStack, len(input))
	for i := range input {
		sources[i] = &LineStack{current: newLine(input[i], 9, i, " ")}
	}

	tree := buildTournamentTree(sources)

	tree = popTopAndUpdate(tree, sources)

}
