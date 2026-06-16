package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func parseBanner(data string) [95][8]string {
	data = strings.ReplaceAll(data, "\r\n", "\n")
	lines := strings.Split(data, "\n")
	var banner [95][8]string
	for i := range 95 {
		start := 1 + i*9
		for j := range 8 {
			if start+j < len(lines) {
				banner[i][j] = lines[start+j]
			}
		}
	}
	return banner
}

func renderArt(input string, banner [95][8]string) string {
	if input == "" {
		return ""
	}
	input = strings.ReplaceAll(input, "\\n", "\n")
	segments := strings.Split(input, "\n")

	var b strings.Builder
	prevEmpty := false

	for _, seg := range segments {
		if seg == "" {
			if !prevEmpty {
				b.WriteByte('\n')
				prevEmpty = true
			}
			continue
		}
		for line := range 8 {
			for _, ch := range seg {
				if ch >= 32 && ch <= 126 {
					b.WriteString(banner[ch-32][line])
				}
			}
			b.WriteByte('\n')
		}
		prevEmpty = false
	}
	return b.String()
}

func parseArgs(args []string) (color, substr, input, banner string, ok bool) {
	if len(args) == 0 {
		return "", "", "", "", false
	}

	if strings.HasPrefix(args[0], "--color=") {
		color = args[0][len("--color="):]
		if color == "" {
			return "", "", "", "", false
		}
		args = args[1:]
	} else if strings.HasPrefix(args[0], "--color") {
		return "", "", "", "", false
	}

	if len(args) == 0 {
		return "", "", "", "", false
	}

	if color != "" {
		switch len(args) {
		case 1:
			return color, "", args[0], "standard", true
		case 2:
			if isKnownBanner(args[1]) {
				return color, "", args[0], args[1], true
			}
			return color, args[0], args[1], "standard", true
		case 3:
			if isKnownBanner(args[2]) {
				return color, args[0], args[1], args[2], true
			}
			return "", "", "", "", false
		}
		return "", "", "", "", false
	}

	switch len(args) {
	case 1:
		return "", "", args[0], "standard", true
	case 2:
		if isKnownBanner(args[1]) {
			return "", "", args[0], args[1], true
		}
		return "", "", "", "", false
	}
	return "", "", "", "", false
}

func isKnownBanner(name string) bool {
	return name == "standard" || name == "shadow" || name == "thinkertoy"
}

func parseColor(c string) string {
	lower := strings.ToLower(c)

	named := map[string]string{
		"black":   "\033[30m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
		"orange":  "\033[38;2;255;165;0m",
		"purple":  "\033[38;2;128;0;128m",
		"pink":    "\033[38;2;255;192;203m",
		"brown":   "\033[38;2;165;42;42m",
	}
	if code, ok := named[lower]; ok {
		return code
	}

	if strings.HasPrefix(lower, "#") {
		r, g, b := parseHex(lower[1:])
		return fmtAnsiTruecolor(r, g, b)
	}

	if strings.HasPrefix(lower, "rgb(") && strings.HasSuffix(lower, ")") {
		inner := lower[4 : len(lower)-1]
		parts := strings.Split(inner, ",")
		if len(parts) == 3 {
			r, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
			g, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
			b, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
			return fmtAnsiTruecolor(r, g, b)
		}
	}

	if strings.HasPrefix(lower, "hsl(") && strings.HasSuffix(lower, ")") {
		inner := lower[4 : len(lower)-1]
		parts := strings.Split(inner, ",")
		if len(parts) == 3 {
			h, _ := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
			sStr := strings.TrimSpace(parts[1])
			s, _ := strconv.ParseFloat(sStr[:len(sStr)-1], 64)
			lStr := strings.TrimSpace(parts[2])
			l, _ := strconv.ParseFloat(lStr[:len(lStr)-1], 64)
			r, g, b := hslToRGB(h, s/100, l/100)
			return fmtAnsiTruecolor(r, g, b)
		}
	}

	if n, err := strconv.Atoi(lower); err == nil && n >= 0 && n <= 255 {
		return fmt.Sprintf("\033[38;5;%dm", n)
	}

	return "\033[0m"
}

func parseHex(hex string) (int, int, int) {
	if len(hex) == 3 {
		r, _ := strconv.ParseInt(string(hex[0])+string(hex[0]), 16, 64)
		g, _ := strconv.ParseInt(string(hex[1])+string(hex[1]), 16, 64)
		b, _ := strconv.ParseInt(string(hex[2])+string(hex[2]), 16, 64)
		return int(r), int(g), int(b)
	}
	if len(hex) >= 6 {
		r, _ := strconv.ParseInt(hex[0:2], 16, 64)
		g, _ := strconv.ParseInt(hex[2:4], 16, 64)
		b, _ := strconv.ParseInt(hex[4:6], 16, 64)
		return int(r), int(g), int(b)
	}
	return 0, 0, 0
}

func hslToRGB(h, s, l float64) (int, int, int) {
	if s == 0 {
		v := int(math.Round(l * 255))
		return v, v, v
	}

	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - c/2

	var r, g, b float64
	switch {
	case h < 60:
		r, g, b = c, x, 0
	case h < 120:
		r, g, b = x, c, 0
	case h < 180:
		r, g, b = 0, c, x
	case h < 240:
		r, g, b = 0, x, c
	case h < 300:
		r, g, b = x, 0, c
	default:
		r, g, b = c, 0, x
	}

	return int(math.Round((r + m) * 255)), int(math.Round((g + m) * 255)), int(math.Round((b + m) * 255))
}

func fmtAnsiTruecolor(r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func buildColorMap(input, substr string) []bool {
	colorMap := make([]bool, len(input))
	if substr == "" {
		for i := range colorMap {
			colorMap[i] = true
		}
		return colorMap
	}
	for i := 0; i <= len(input)-len(substr); i++ {
		if input[i:i+len(substr)] == substr {
			for j := 0; j < len(substr); j++ {
				colorMap[i+j] = true
			}
		}
	}
	return colorMap
}

func renderArtColor(input string, banner [95][8]string, colorCode, substr string) string {
	if input == "" {
		return ""
	}
	input = strings.ReplaceAll(input, "\\n", "\n")
	segments := strings.Split(input, "\n")

	resetCode := "\033[0m"
	var b strings.Builder
	prevEmpty := false

	for _, seg := range segments {
		if seg == "" {
			if !prevEmpty {
				b.WriteByte('\n')
				prevEmpty = true
			}
			continue
		}

		colorMap := buildColorMap(seg, substr)

		for line := range 8 {
			inColor := false
			for i, ch := range seg {
				if ch >= 32 && ch <= 126 {
					if colorMap[i] && !inColor {
						b.WriteString(colorCode)
						inColor = true
					} else if !colorMap[i] && inColor {
						b.WriteString(resetCode)
						inColor = false
					}
					b.WriteString(banner[ch-32][line])
				}
			}
			if inColor {
				b.WriteString(resetCode)
			}
			b.WriteByte('\n')
		}
		prevEmpty = false
	}
	return b.String()
}
