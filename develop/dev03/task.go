package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Config
const (
	bufferSize = 10 // line is 4096 bytes at worst since max length of line in file is 2048 bytes
)

// Arguments
var (
	filePath string
	KEY      int
)

type HeapElem struct { // utility struct
	scanner *bufio.Scanner
	line    *line
}

func init() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	// Remove old chunks
	files, err := filepath.Glob(".chunk*")
	if err != nil {
		log.Fatalln(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			log.Fatalln(err)
		}
	}

	flag.IntVar(&KEY, "k", 9,
		"Position of column used for sorting. Lines are separated into columns by spacebar.",
	)
}

func main() {
	flag.Parse() // need to be called from inside `main()`, otherwise tests fail

	if l := len(os.Args); l > 0 {
		filePath = os.Args[l-1] // path to file is always the last argument
	} else {
		log.Fatalln("no file to sort")
	}

	filePath = "input.log"

	in, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	// Scan and sort strings from file into chunks.
	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)
	buf := make([]*line, bufferSize)[0:0] // set capacity to 0 to avoid directly indexing buffer
	chunks := []string{}
	for s.Scan() {
		if len(buf) < bufferSize {
			buf = append(buf, newLine(s.Text(), KEY, " "))
		} else {
			mergeSort(buf)
			chunks = append(chunks, writeToFile(buf)) // storing names of chunks for later use
			buf = buf[0:0]                            // emptying buffer on write out
		}
	}
	if len(buf) != 0 { // flushing buffer if there is anything left in it
		mergeSort(buf)
		chunks = append(chunks, writeToFile(buf))
	}
	// defer deleteChunks(chunks)

	if s.Err() != nil {
		log.Fatalln(s.Err())
	}

	// Creating an array for tree for mergin chunks
	tree := make([]*HeapElem, len(chunks))
	for i := 0; i < len(chunks); i++ {
		f, err := os.Open(chunks[i])
		if err != nil {
			log.Fatalln(err)
		}
		tree[i] = &HeapElem{}
		tree[i].scanner = bufio.NewScanner(f)
		tree[i].scanner.Split(bufio.ScanLines)
		if tree[i].scanner.Scan() {
			tree[i].line = newLine(tree[i].scanner.Text(), KEY, " ")
		}
	}

	// Merging chunks
	out, err := os.OpenFile("out", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	tree = buildTournamentTree(tree) // constructing tournament tree to speed up merging
	for len(tree) != 1 {
		prev := tree[0].line
		out.WriteString(tree[0].line.line + "\n")
		if tree[0].scanner.Scan() {
			tree[0].line = newLine(tree[0].scanner.Text(), KEY, " ")
			tree = repairTournamentTree(tree)
		} else {
			tree = tree[1:]
		}
		if prev.bigger(tree[0].line) {
			log.Panicf("%s\n%s\n",
				tree[0].line.line, prev.line,
			)
		}
	}

	for tree[0].scanner.Scan() {
		out.WriteString(tree[0].scanner.Text() + "\n")
	}
}

func buildTournamentTree(tree []*HeapElem) []*HeapElem {
	cart, stage := 0, 1
	for stage <= len(tree)/2+len(tree)%2 { // %2 for when last index is non-odd
		for cart+stage < len(tree) {
			a := tree[cart].line
			b := tree[cart+stage].line
			if a.bigger(b) { // TODO: smaller ???
				tree[cart], tree[cart+stage] = tree[cart+stage], tree[cart]
			}

			cart += stage * 2
		}
		stage++
		cart = 0
	}

	return tree
}

func repairTournamentTree(tree []*HeapElem) []*HeapElem {
	cart, stage := 0, 1
	for stage <= len(tree)/2 {
		for cart+stage < len(tree) {
			a := tree[0].line
			b := tree[cart+stage].line

			if a.getKey() == "9065}" || b.getKey() == "9065}" {
				fmt.Println("Gotcha")
			}

			if a.bigger(b) {
				tree[0], tree[cart+stage] = tree[cart+stage], tree[0]
			}

			cart += stage * 2
		}
		stage++
		cart = 0
	}

	return tree
}

func writeToFile(arr []*line) (fileName string) {
	file, err := os.CreateTemp(".", ".chunk*")
	if err != nil {
		log.Fatalf("failed to create temporary file: %s\n", err)
	}
	defer file.Close()

	for _, l := range arr {
		file.Write([]byte(l.line + "\n"))
	}

	return file.Name()
}

func deleteChunks(arr []string) {
	for _, f := range arr {
		os.Remove(f)
	}
}
