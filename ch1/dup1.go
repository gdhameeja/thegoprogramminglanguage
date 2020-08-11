package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Println(line, count)
		}
	}
}

