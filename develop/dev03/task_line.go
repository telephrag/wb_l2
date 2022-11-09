package main

import "strings"

type LineCompEnum int

const (
	lineEquals  LineCompEnum = 0
	lineBigger  LineCompEnum = 1
	lineSmaller LineCompEnum = 2
)

type line struct {
	keyStart int    // index at which key column starts, -1 means there is no key
	keyEnd   int    // index at which key column ends
	line     string // line from a file
}

// Constructs a new `line` object.
// `str` -- the whole line
// `key` -- position of sorting key, starts at 0
// `sep` -- separator of words in line, key will be a word seperated by it
func newLine(str string, key int, sep string) *line {
	l := &line{}
	l.line = str
	words := strings.Split(str, sep)
	if len(words) > key {
		l.keyStart += key
		var i int
		for i = 0; i < key; i++ {
			l.keyStart += len(words[i])
		}
		l.keyEnd = l.keyStart + len(words[i])
	} else {
		l.keyStart = -1
		l.keyEnd = -1
	}
	return l
}

func (l *line) hasKey() bool {
	return l.keyStart != -1
}

func (l *line) getKey() string {
	if !l.hasKey() {
		return ""
	}
	return l.line[l.keyStart:l.keyEnd]
}

// Returns `true` if `l` is lexicographically bigger than `other`, `false` otherwise.
func (l *line) compare(other *line) LineCompEnum {

	lHasKey := l.hasKey()
	otherHasKey := other.hasKey()

	if lHasKey && otherHasKey {
		switch {
		case l.getKey() > other.getKey():
			return lineBigger
		case l.getKey() < other.getKey():
			return lineSmaller
		default:
			return lineEquals
		}
	}

	if lHasKey && !otherHasKey {
		return lineBigger
	}

	if !lHasKey && otherHasKey {
		return lineSmaller
	}

	switch {
	case l.line > other.line:
		return lineBigger
	case l.line < other.line:
		return lineSmaller
	default:
		return lineEquals
	}
}

func (l *line) bigger(other *line) bool { return l.compare(other) == lineBigger }

func (l *line) smaller(other *line) bool { return l.compare(other) == lineSmaller }

func (l *line) biggerOrEquals(other *line) bool {
	res := l.compare(other)
	return res == lineBigger || res == lineEquals
}

func (l *line) smallerOrEquals(other *line) bool {
	res := l.compare(other)
	return res == lineSmaller || res == lineEquals
}

func (l *line) equals(other *line) bool { return l.compare(other) == lineEquals }
