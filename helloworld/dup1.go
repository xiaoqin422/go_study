package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}
	fmt.Printf("%v\n", counts)
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
