package main

import (
	"bufio"
	"os"
)

func getPosition(char byte) int {
	return (int(char-32) * 855 / 95) + 2
}

func getLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

func getCharLines(char byte, lines []string) []string {
	out := make([]string, 0)

	startIndex := getPosition(char)
	for i := startIndex - 1; i < startIndex+7; i++ {
		out = append(out, lines[i])
	}
	return out
}

func getStringLines(s string, lines []string) [][]string {
	out := make([][]string, 0)

	for _, b := range s {
		out = append(out, getCharLines(byte(b), lines))
	}
	return out
}
