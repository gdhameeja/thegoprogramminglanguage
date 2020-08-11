package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, arg := range os.Args[1:] {
		content, err := ioutil.ReadFile(arg)
		if err != nil {
			fmt.Println("Something went wrong while reading the file")
		}
		countlines(string(content), counts)
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n", line, n)
		}
	}
}

func countlines(content string, counts map[string]int) {
	clines := strings.Split(content, "\n")
	for _, line := range clines {
		counts[line]++
	}
}
