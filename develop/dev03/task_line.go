package develop

import "strings"

type LineCompEnum int

const (
	lineEquals  LineCompEnum = 0
	lineBigger  LineCompEnum = 1
	lineSmaller LineCompEnum = 2
)

type line struct {
	key  string
	line string
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
		l.key = words[key]
	}
	return l
}

// Returns `true` if `l` is lexicographically bigger than `other`, `false` otherwise.
func (l *line) compare(other *line) LineCompEnum {

	hasKey := l.key != ""
	otherHasKey := other.key != ""

	if hasKey && otherHasKey {
		switch {
		case l.key > other.key:
			return lineBigger
		case l.key < other.key:
			return lineSmaller
		default:
			return lineEquals
		}
	}

	if hasKey && !otherHasKey {
		return lineBigger
	}

	if !hasKey && otherHasKey {
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
