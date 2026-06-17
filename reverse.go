package main

import (
	"io/fs"
	"os"
	"strings"
)

func loadBanners() [][95][8]string {
	var result [][95][8]string
	for _, name := range []string{"standard", "shadow", "thinkertoy"} {
		data, err := fs.ReadFile(os.DirFS("."), name+".txt")
		if err != nil {
			continue
		}
		result = append(result, parseBanner(string(data)))
	}
	return result
}

func reverseArt(data string, banners ...[95][8]string) string {
	data = normalizeArt(data)
	if data == "" {
		return ""
	}

	for _, banner := range banners {
		result := decodeWithBanner(data, banner)

		reEncoded := renderArt(result, banner)
		if normalizeArt(reEncoded) == data {
			return result
		}
	}

	var best string
	for _, banner := range banners {
		result := decodeWithBanner(data, banner)
		if len(result) > len(best) {
			best = result
		}
	}
	return best
}

func normalizeArt(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	return strings.TrimRight(s, "\n")
}

func decodeWithBanner(data string, banner [95][8]string) string {
	lines := strings.Split(data, "\n")

	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	if len(lines) == 0 {
		return ""
	}

	var result []string
	var group []string

	for _, line := range lines {
		if line == "" {
			if len(group) == 8 {
				result = append(result, decodeArtGroup(group, banner))
			}
			group = nil
			result = append(result, "")
			continue
		}
		group = append(group, line)
		if len(group) == 8 {
			result = append(result, decodeArtGroup(group, banner))
			group = nil
		}
	}

	if len(group) == 8 {
		result = append(result, decodeArtGroup(group, banner))
	}

	return strings.Join(result, "\n")
}

func decodeArtGroup(lines []string, banner [95][8]string) string {
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}
	padded := make([]string, 8)
	for i, line := range lines {
		padded[i] = line + strings.Repeat(" ", maxWidth-len(line))
	}

	var out strings.Builder
	pos := 0
	for pos < maxWidth {
		matched := false
		for c := 94; c >= 0; c-- {
			w := len(banner[c][0])
			if pos+w > maxWidth {
				continue
			}
			ok := true
			for row := 0; row < 8; row++ {
				if padded[row][pos:pos+w] != banner[c][row] {
					ok = false
					break
				}
			}
			if ok {
				if c == 0 && allSpaces(padded, pos+w, maxWidth) {
					pos = maxWidth
					matched = true
					break
				}
				out.WriteRune(rune(32 + c))
				pos += w
				matched = true
				break
			}
		}
		if !matched {
			break
		}
	}
	return out.String()
}

func allSpaces(padded []string, start, end int) bool {
	for _, line := range padded {
		for _, ch := range line[start:end] {
			if ch != ' ' {
				return false
			}
		}
	}
	return true
}
