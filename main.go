package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	color, substr, align, input, banner, ok := parseArgs(os.Args[1:])
	if !ok {
		printUsage()
		return
	}

	bannerFile := banner + ".txt"
	data, err := fs.ReadFile(os.DirFS("."), bannerFile)
	if err != nil {
		printUsage()
		return
	}

	bannerData := parseBanner(string(data))

	termWidth := getTerminalWidth()

	if color != "" {
		colorCode := parseColor(color)
		result := renderArtColor(input, bannerData, colorCode, substr)
		fmt.Print(result)
	} else {
		result := renderArtAligned(input, bannerData, align, termWidth)
		fmt.Print(result)
	}
}

func printUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . --align=right something standard")
}
