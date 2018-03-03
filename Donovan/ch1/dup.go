package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		counts["stdin"] = make(map[string]int)
		countLines(os.Stdin, counts["stdin"])
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			counts[arg] = make(map[string]int)
			countLines(f, counts[arg])
			f.Close()
		}
	}
	for fname, count := range counts {
		for line, n := range count {
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, fname)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
