package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

// Config
const (
	bufferSize = 10 // line is 4096 bytes at worst since max length of line in file is 2048 bytes
)

// Arguments
var (
	filePath string
	key      int
)

func init() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	flag.IntVar(&key, "k", 0,
		"Position of column used for sorting. Lines are separated into columns by spacebar.",
	)

}

func main() {
	flag.Parse() // need to be called from inside `main()`, otherwise tests fail

	if l := len(os.Args); l > 0 {
		filePath = os.Args[l-1]
	} else {
		log.Fatalln("no file to sort")
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	buf := make([]*line, bufferSize)[0:0]

	for s.Scan() {
		if len(buf) < bufferSize {
			buf = append(buf, newLine(s.Text(), key, " "))
		} else {
			mergeSort(buf)
			writeToFile(buf)
			buf = buf[0:0]
		}
	}

}

func writeToFile(arr []*line) {
	file, err := os.CreateTemp(".", ".chunk*")
	if err != nil {
		log.Fatalf("failed to create temporary file: %s\n", err)
	}

	for _, l := range arr {
		file.Write([]byte(l.line + "\n"))
	}
}
