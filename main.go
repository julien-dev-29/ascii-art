package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input := ""
	banner := "standard"
	color := ""
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Println("Usage: go run main.go [input] [banner]")
			fmt.Println("Example: go run main.go \"Hello World\" shadow")
			return
		}
		if after, ok := strings.CutPrefix(os.Args[1], "--color="); ok {
			color = after
		}
		// Si il y a une couleur spécifiée, on la retire des arguments
		if color != "" {
			os.Args = append(os.Args[:1], os.Args[2:]...)
		}
		//Si il y a un substring de la chaine de caractères a mettre en couleur, on le retire des arguments
		input = os.Args[1]
		if len(os.Args) > 2 {
			banner = os.Args[2]
		}
	}

	input = strings.ReplaceAll(input, "\\n", "\n")

	m := parseFile(banner + ".txt")

	if input == "" {
		return
	}

	var b strings.Builder
	segments := strings.SplitSeq(input, "\n")
	for seg := range segments {
		if seg == "" {
			b.WriteByte('\n')
			continue
		}
		for line := range 8 {
			for _, ch := range seg {
				if lines, ok := m[rune(ch)]; ok {
					b.WriteString(lines[line])
				}
			}
			b.WriteByte('\n')
		}
	}
	// Display the ASCII art with color if specified
	if colorCode := getColorCode(color); colorCode != "" {
		fmt.Print(colorCode)
	}
	fmt.Print(b.String())
	if colorCode := getColorCode(color); colorCode != "" {
		fmt.Print("\033[0m") // Reset color
	}
}
