package main

import "log"

func getSmaller(first, second *Line) *Line {
	if first.Smaller(second) {
		return first
	}
	return second
}

// sources must have odd number of elements
func buildTournamentTree(sources []*LineStack) []*Line {
	if len(sources) < 2 {
		return []*Line{sources[0].Get(), sources[0].Get(), &LINE_INFINITY}
	}

	l := len(sources)
	tree := make([]*Line, l-1)
	this := tree[len(tree)-l/2-l%2:]
	var i, j int
	for i, j = 0, 0; i < l-l%2; i += 2 { // <= instead of < ???
		this[j] = getSmaller(
			sources[i].Get(),
			sources[i+1].Get(),
		)
		j++
	}
	if l%2 != 0 { // supplement pairless line with `LINE_INFINITY`
		this[j] = sources[i].Get()
		j++
	}

	cart := len(tree) - j // j ~ len(this) here

	// DEBUG
	// for i := range this {
	// 	fmt.Print(this[i].getKey())
	// }
	// fmt.Println()

	next := tree[cart-j/2-j%2 : cart]
	for len(this) != 1 {
		for i := 0; i*2+1 < len(this); i++ {
			next[i] = getSmaller(this[i*2], this[i*2+1])
		}
		if len(this)%2 != 0 {
			next[len(next)-1] = this[len(this)-1]
		}
		cart -= len(next)
		this = next
		l = len(this)
		next = tree[cart-l/2 : cart]

		// // DEBUG
		// for i := range this {
		// 	fmt.Print(this[i].getKey())
		// }
		// fmt.Println()
	}

	return tree
}

func popTopAndUpdate(tree []*Line, sources []*LineStack) []*Line {
	path := []int{0}

	last := func([]int) int {
		return path[len(path)-1]
	}

	id := tree[0].sourceID
	lc := levelsCount(tree)
	for i := 1; i < lc; i++ {
		if tree[2*last(path)+1].sourceID == id {
			path = append(path, 2*last(path)+1)
		} else if tree[2*last(path)+2].sourceID == id {
			path = append(path, 2*last(path)+2)
		} else {
			log.Fatalln("invalid tree, iteration: ", i)
		}
	}

	// // DEBUG
	// elem := newLine(
	// 	"{\"thread\": 5, \"timestamp\": 2022-11-08 13:50:17.461054186 +0400 +04 m=+0.006923689, \"result\": 0}",
	// 	9, -1, " ",
	// )
	// id := tree[last(path)].sourceID // same to aquired previously ???
	elem := sources[id].Pop()
	tree[last(path)] = elem
	// if odd amount of elements left to elem than we should compare elem
	// to element to his right, otherwise compare with element to it's left
	for i := len(path) - 2; i >= 0; i-- {
		p := path[i]
		tree[p] = getSmaller(tree[2*p+1], tree[2*p+2])
	}

	return tree
}

func levelsCount(tree []*Line) int {
	ln := len(tree)
	res := 0
	for ln != 0 {
		ln /= 2
		res++
	}
	return res
}
