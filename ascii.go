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

func getALines() {

}
