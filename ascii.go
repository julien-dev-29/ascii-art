package main

import (
	"strings"
)

func getLetterByCodeAscci() {

}

func splitByLine(data []byte) []string {
	out := make([]string, 0)
	for line := range strings.Lines(string(data)) {
		out = append(out, line)
	}
	return out
}
