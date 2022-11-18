package main

import (
	"bufio"
	"flag"
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
	buf := make([]*Line, bufferSize)[0:0] // set capacity to 0 to avoid directly indexing buffer
	chunks := []string{}
	for s.Scan() {
		if len(buf) < bufferSize {
			buf = append(buf, newLine(s.Text(), KEY, -1, " "))
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

	// Creating an array for sources for mergin chunks
	sources := make([]*LineStack, len(chunks))
	for i := 0; i < len(chunks); i++ {
		f, err := os.Open(chunks[i])
		if err != nil {
			log.Fatalln(err)
		}
		sources[i] = (&LineStack{}).Init(f, KEY, i)
		sources[i].Pop()
	}

	// Merging chunks
	out, err := os.OpenFile("out", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	tree := buildTournamentTree(sources)
	for !tree[0].Same(&LINE_INFINITY) {
		out.WriteString(tree[0].line + "\n")
		tree = popTopAndUpdate(tree, sources)
	}
}

func writeToFile(arr []*Line) (fileName string) {
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
