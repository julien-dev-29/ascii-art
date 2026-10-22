package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must provide a string")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too much arguments")
		return
	}
	s := os.Args[1]
	lines := getLines("standard.txt")
	strLines := strings.Builder{}
	if containAntiSlashN(s) {
		words := strings.SplitSeq(s, "\n")
		for word := range words {
			getStringLines(word, lines)
		}
	} else {
		for i := range 8 {
			for _, l := range getStringLines(s, lines) {
				strLines.WriteString(l[i])
			}
			strLines.WriteRune('\n')
		}
	}
	fmt.Println(strLines.String())
}
