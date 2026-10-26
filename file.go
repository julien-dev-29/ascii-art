package main

import (
	"bufio"
	"os"
)

func parseFile(path string) map[rune][]string {
	result := make(map[rune][]string)
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())

	char := rune(32)

	if len(lines) > 0 && len(lines)%9 == 0 && lines[0] == "" {
		for i := 0; i < len(lines)/9; i++ {
			start := 1 + i*9
			if start+8 > len(lines) {
				break
			}
			result[char] = lines[start : start+8]
			char++
		}
	} else {
		for i := 0; i+8 <= len(lines); i += 8 {
			result[char] = lines[i : i+8]
			char++
		}
	}

	return result
}
