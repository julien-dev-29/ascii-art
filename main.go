package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	input := os.Args[1]
	data, _ := os.ReadFile("standard.txt")
	lines := strings.Split(string(data), "\n")
	for i := range 8 {
		for _, c := range input {
			if c == '\n' {
				fmt.Println()
				continue
			}
			index := int(c - ' ')
			start := index * 9
			fmt.Print(lines[start+i])
		}
		fmt.Println()
	}
}
