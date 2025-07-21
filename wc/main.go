package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	filename := flag.String("file", "", "Read from file")
	flag.Parse()

	reader := os.Stdin

	if *filename != "" {
		file, err := os.Open(*filename)
		if err != nil {
			fmt.Println("Couldn't read file")
			os.Exit(1)
		}
		defer file.Close()

		reader = file
	}

	fmt.Println(count(reader, *lines, *bytes))
}

func count(r io.Reader, countLines, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	if countBytes && !countLines {
		scanner.Split(bufio.ScanBytes)
	}

	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
