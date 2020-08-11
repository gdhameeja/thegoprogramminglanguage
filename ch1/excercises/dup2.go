package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]string)
	if len(os.Args[1:]) == 0 {
		countlines(os.Stdin, "Standard Input", counts, filenames)
	} else {
		for _, filepath := range os.Args[1:] {
			file, err := os.Open(filepath); if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: something went wrong while reading the file")
				continue
			}
			countlines(file, filepath, counts, filenames)
		}
	}
	
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%s\t%d\n", line, filenames[line], n)
		}
	}
}

func countlines(file *os.File, filename string, counts map[string]int, filenames map[string]string) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp := scanner.Text()
		counts[temp]++
		filenames[temp] = filename

	}
}
