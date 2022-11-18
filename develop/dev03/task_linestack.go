package main

import (
	"bufio"
	"os"
)

var LINE_INFINITY = Line{-1, 0, 8, "INFINITY", true}

type LineStack struct {
	id      int
	scanner *bufio.Scanner
	current *Line
	key     int
}

func (s *LineStack) Init(f *os.File, key, id int) (self *LineStack) {
	s.id = id
	s.scanner = bufio.NewScanner(f)
	s.scanner.Split(bufio.ScanLines)
	s.key = key
	return s
}

func (s *LineStack) Pop() *Line {
	if s.scanner.Scan() {
		s.current = newLine(s.scanner.Text(), s.key, s.id, " ")
		return s.current
	} else {
		s.current = &LINE_INFINITY
		return &LINE_INFINITY
	}
}

func (s *LineStack) Get() *Line {
	return s.current
}
