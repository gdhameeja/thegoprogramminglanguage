// this dup program takes in files as well. If no files is passed as command line arguement
// Stdin is used as input
package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	if len(os.Args[1:]) == 0 {
		countlines(os.Stdin, counts)
	} else {
		for _, filepath := range os.Args[1:] {
			filepointer, err := os.Open(filepath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countlines(filepointer, counts)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n", line, n)
		}
	}
}

func countlines(filepointer *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(filepointer)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
}

