package main

import (
	"strings"
)

type LineCompEnum int

const (
	lineEquals  LineCompEnum = 0
	lineBigger  LineCompEnum = 1
	lineSmaller LineCompEnum = 2
)

type Line struct {
	sourceID   int
	keyStart   int    // index at which key column starts, -1 means there is no key
	keyEnd     int    // index at which key column ends
	line       string // line from a file
	isInfinity bool
}

// Constructs a new `line` object.
// `str` -- the whole line
// `key` -- position of sorting key, starts at 0
// `sep` -- separator of words in line, key will be a word seperated by it
func newLine(str string, key, sourceID int, sep string) *Line {
	l := &Line{}
	l.line = str
	l.sourceID = sourceID
	words := strings.Split(str, sep)
	if len(words) > key {
		l.keyStart += key // adding amount of spacebars to keyStart
		var i int
		for i = 0; i < key; i++ {
			l.keyStart += len(words[i]) // adding length of all words before key column
		}
		l.keyEnd = l.keyStart + len(words[i]) // adding length of key column itself
	} else {
		l.keyStart = -1
		l.keyEnd = -1
	}
	return l
}

func (l *Line) HasKey() bool {
	return l.keyStart != -1
}

func (l *Line) GetKey() string {
	if !l.HasKey() {
		return ""
	}
	return l.line[l.keyStart:l.keyEnd]
}

// Returns `true` if `l` is lexicographically bigger than `other`, `false` otherwise.
func (l *Line) Compate(other *Line) LineCompEnum {

	thisInf := l.isInfinity
	otherInf := other.isInfinity
	if thisInf || otherInf {
		switch {
		case thisInf && otherInf:
			return lineEquals
		case thisInf:
			return lineBigger
		case otherInf:
			return lineSmaller
		}
	}

	thisHasKey := l.HasKey()
	otherHasKey := other.HasKey()

	if thisHasKey && otherHasKey {
		switch {
		case l.GetKey() > other.GetKey():
			return lineBigger
		case l.GetKey() < other.GetKey():
			return lineSmaller
		default:
			return lineEquals
		}
	}

	if thisHasKey && !otherHasKey {
		return lineBigger
	}

	if !thisHasKey && otherHasKey {
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

func (l *Line) Bigger(other *Line) bool { return l.Compate(other) == lineBigger }

func (l *Line) Smaller(other *Line) bool { return l.Compate(other) == lineSmaller }

func (l *Line) BiggerOrEquals(other *Line) bool {
	res := l.Compate(other)
	return res == lineBigger || res == lineEquals
}

func (l *Line) SmallerOrEquals(other *Line) bool {
	res := l.Compate(other)
	return res == lineSmaller || res == lineEquals
}

func (l *Line) Equals(other *Line) bool { return l.Compate(other) == lineEquals }

func (l *Line) Same(other *Line) bool {
	if l.isInfinity && other.isInfinity {
		return true
	}

	if l.keyStart != other.keyStart {
		return false
	}

	if l.keyEnd != other.keyEnd {
		return false
	}

	if l.line != other.line {
		return false
	}

	return true
}
