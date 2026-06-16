package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	color, substr, input, banner, ok := parseArgs(os.Args[1:])
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

	if color != "" {
		colorCode := parseColor(color)
		result := renderArtColor(input, bannerData, colorCode, substr)
		fmt.Print(result)
	} else {
		result := renderArt(input, bannerData)
		fmt.Print(result)
	}
}

func printUsage() {
	fmt.Println("Usage: go run . [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . something standard")
}
