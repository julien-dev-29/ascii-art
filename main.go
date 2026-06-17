package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	arg := os.Args[1]
	if strings.HasPrefix(arg, "--reverse=") {
		fileName := arg[len("--reverse="):]
		if fileName == "" {
			printUsage()
			return
		}
		data, err := fs.ReadFile(os.DirFS("."), fileName)
		if err != nil {
			printUsage()
			return
		}
		banners := loadBanners()
		result := reverseArt(string(data), banners...)
		fmt.Print(result)
		return
	}

	color, substr, align, input, banner, outputFile, ok := parseArgs(os.Args[1:])
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

	var result string
	if color != "" {
		colorCode := parseColor(color)
		result = renderArtColor(input, bannerData, colorCode, substr)
	} else {
		result = renderArtAligned(input, bannerData, align, termWidth)
	}

	if outputFile != "" {
		os.WriteFile(outputFile, []byte(result), 0644)
	} else {
		fmt.Print(result)
	}
}

func printUsage() {
	fmt.Println("Usage: go run . [OPTION]")
	fmt.Println()
	fmt.Println("EX: go run . --reverse=<fileName>")
}
