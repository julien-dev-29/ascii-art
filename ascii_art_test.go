package main

import (
	"os"
	"strings"
	"testing"
)

func loadStdTestBanner(t *testing.T) [95][8]string {
	t.Helper()
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		t.Fatal("failed to load standard.txt:", err)
	}
	return parseBanner(string(data))
}

func TestParseBanner_Has95Chars(t *testing.T) {
	banner := loadStdTestBanner(t)
	if len(banner) != 95 {
		t.Errorf("expected 95 chars, got %d", len(banner))
	}
}

func TestParseBanner_FirstCharIsSpace(t *testing.T) {
	banner := loadStdTestBanner(t)
	for _, line := range banner[0] {
		if strings.TrimSpace(line) != "" {
			t.Errorf("expected space char (index 0) to be all spaces, got %q", line)
		}
	}
}

func TestParseBanner_ExclamationMark(t *testing.T) {
	banner := loadStdTestBanner(t)
	want := " _  "
	got := banner[1][0]
	if got != want {
		t.Errorf("expected first line of '!' to be %q, got %q", want, got)
	}
}

func TestRenderArt_EmptyInput(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArt("", banner)
	if got != "" {
		t.Errorf("expected empty string, got %q", got)
	}
}

func TestRenderArt_SingleNewline(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArt("\\n", banner)
	want := "\n"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestRenderArt_DoubleNewlineCollapsed(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArt("\\n\\n", banner)
	lines := strings.Split(strings.TrimRight(got, "\n"), "\n")
	if len(lines) != 1 {
		t.Errorf("expected 1 blank line, got %d lines", len(lines))
	}
}

func TestRenderArt_NewlineBetweenText(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArt("\\n", banner)
	want := "\n"
	if got != want {
		t.Errorf("expected single newline, got %q", got)
	}
}

func TestRenderArt_Hello(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArt("Hello", banner)
	n := strings.Count(got, "\n")
	if n != 8 {
		t.Errorf("expected 8 newlines for 'Hello', got %d", n)
	}
}

func TestRenderArt_HelloNewlineThere(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArt("Hello\\nThere", banner)
	n := strings.Count(got, "\n")
	if n != 16 {
		t.Errorf("expected 16 newlines for 'Hello\\nThere', got %d", n)
	}
}

func TestRenderArt_HelloDoubleNewlineThere(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArt("Hello\\n\\nThere", banner)
	n := strings.Count(got, "\n")
	if n != 17 {
		t.Errorf("expected 17 newlines for 'Hello\\n\\nThere', got %d", n)
	}
}

func TestRenderArt_TrailingNewline(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArt("Hello\\n", banner)
	n := strings.Count(got, "\n")
	if n != 9 {
		t.Errorf("expected 9 newlines for 'Hello\\n', got %d", n)
	}
}

func TestRenderArt_NonPrintableCharsAreSkipped(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArt("A\x00B", banner)
	want := renderArt("AB", banner)
	if got != want {
		t.Errorf("expected non-printable chars to be skipped, got:\n%q\nwant:\n%q", got, want)
	}
}

func TestRenderArt_SpecialCharacters(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArt("{Hello There}", banner)
	n := strings.Count(got, "\n")
	if n != 8 {
		t.Errorf("expected 8 newlines, got %d", n)
	}
}

func TestRenderArt_Numbers(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArt("123", banner)
	n := strings.Count(got, "\n")
	if n != 8 {
		t.Errorf("expected 8 newlines, got %d", n)
	}
}

func TestRenderArt_AllBannersSameHeight(t *testing.T) {
	standard := loadStdTestBanner(t)
	shadowData, err := os.ReadFile("shadow.txt")
	if err != nil {
		t.Fatal(err)
	}
	shadow := parseBanner(string(shadowData))

	thinkertoyData, err := os.ReadFile("thinkertoy.txt")
	if err != nil {
		t.Fatal(err)
	}
	thinkertoy := parseBanner(string(thinkertoyData))

	gotStd := renderArt("Hello", standard)
	gotShadow := renderArt("Hello", shadow)
	gotThink := renderArt("Hello", thinkertoy)

	if strings.Count(gotStd, "\n") != 8 {
		t.Errorf("standard: expected 8 newlines, got %d", strings.Count(gotStd, "\n"))
	}
	if strings.Count(gotShadow, "\n") != 8 {
		t.Errorf("shadow: expected 8 newlines, got %d", strings.Count(gotShadow, "\n"))
	}
	if strings.Count(gotThink, "\n") != 8 {
		t.Errorf("thinkertoy: expected 8 newlines, got %d", strings.Count(gotThink, "\n"))
	}
}

func TestRenderArt_MultipleNewlinesCollapsed(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArt("A\\n\\n\\n\\nB", banner)
	n := strings.Count(got, "\n")
	if n != 17 {
		t.Errorf("expected 17 newlines (8+1+8), got %d", n)
	}
}

func TestRenderArt_InputHasOnlySpaces(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArt("   ", banner)
	lines := strings.Split(strings.TrimRight(got, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("expected 8 lines for spaces, got %d", len(lines))
	}
	for _, line := range lines {
		if strings.TrimRight(line, " ") != "" {
			t.Errorf("expected only spaces, got %q", line)
		}
	}
}

func TestParseColor_NamedRed(t *testing.T) {
	got := parseColor("red")
	want := "\033[31m"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestParseColor_NamedBlue(t *testing.T) {
	got := parseColor("BLUE")
	want := "\033[34m"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestParseColor_Hex(t *testing.T) {
	got := parseColor("#FF0000")
	want := "\033[38;2;255;0;0m"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestParseColor_ShortHex(t *testing.T) {
	got := parseColor("#f00")
	want := "\033[38;2;255;0;0m"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestParseColor_RGB(t *testing.T) {
	got := parseColor("rgb(0,255,0)")
	want := "\033[38;2;0;255;0m"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestParseColor_RGBWithSpaces(t *testing.T) {
	got := parseColor("rgb(0, 255, 0)")
	want := "\033[38;2;0;255;0m"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestParseColor_HSL(t *testing.T) {
	got := parseColor("hsl(0,100%,50%)")
	want := "\033[38;2;255;0;0m"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestParseColor_ANSI256(t *testing.T) {
	got := parseColor("196")
	want := "\033[38;5;196m"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestParseColor_Unrecognized(t *testing.T) {
	got := parseColor("notacolor")
	want := "\033[0m"
	if got != want {
		t.Errorf("expected reset code %q, got %q", want, got)
	}
}

func TestParseColor_Orange(t *testing.T) {
	got := parseColor("orange")
	want := "\033[38;2;255;165;0m"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestParseHex_6Digit(t *testing.T) {
	r, g, b := parseHex("FF0000")
	if r != 255 || g != 0 || b != 0 {
		t.Errorf("expected (255,0,0), got (%d,%d,%d)", r, g, b)
	}
}

func TestParseHex_3Digit(t *testing.T) {
	r, g, b := parseHex("f00")
	if r != 255 || g != 0 || b != 0 {
		t.Errorf("expected (255,0,0), got (%d,%d,%d)", r, g, b)
	}
}

func TestParseHex_Invalid(t *testing.T) {
	r, g, b := parseHex("")
	if r != 0 || g != 0 || b != 0 {
		t.Errorf("expected (0,0,0), got (%d,%d,%d)", r, g, b)
	}
}

func TestHSLToRGB_Red(t *testing.T) {
	r, g, b := hslToRGB(0, 1, 0.5)
	if r != 255 || g != 0 || b != 0 {
		t.Errorf("expected (255,0,0), got (%d,%d,%d)", r, g, b)
	}
}

func TestHSLToRGB_Green(t *testing.T) {
	r, g, b := hslToRGB(120, 1, 0.5)
	if r != 0 || g != 255 || b != 0 {
		t.Errorf("expected (0,255,0), got (%d,%d,%d)", r, g, b)
	}
}

func TestHSLToRGB_Blue(t *testing.T) {
	r, g, b := hslToRGB(240, 1, 0.5)
	if r != 0 || g != 0 || b != 255 {
		t.Errorf("expected (0,0,255), got (%d,%d,%d)", r, g, b)
	}
}

func TestHSLToRGB_Gray(t *testing.T) {
	r, g, b := hslToRGB(0, 0, 0.5)
	if r != 128 || g != 128 || b != 128 {
		t.Errorf("expected (128,128,128), got (%d,%d,%d)", r, g, b)
	}
}

func TestBuildColorMap_EmptySubstr(t *testing.T) {
	cm := buildColorMap("hello", "")
	expected := []bool{true, true, true, true, true}
	for i, v := range cm {
		if v != expected[i] {
			t.Errorf("position %d: expected %v, got %v", i, expected[i], v)
		}
	}
}

func TestBuildColorMap_NoMatch(t *testing.T) {
	cm := buildColorMap("hello", "z")
	for i, v := range cm {
		if v {
			t.Errorf("position %d: expected false, got true", i)
		}
	}
}

func TestBuildColorMap_SingleMatch(t *testing.T) {
	cm := buildColorMap("hello", "ll")
	expected := []bool{false, false, true, true, false}
	for i, v := range cm {
		if v != expected[i] {
			t.Errorf("position %d: expected %v, got %v", i, expected[i], v)
		}
	}
}

func TestBuildColorMap_MultipleMatches(t *testing.T) {
	cm := buildColorMap("a king kitten have kit", "kit")
	if !cm[7] || !cm[8] || !cm[9] {
		t.Errorf("expected 'kit' starting at pos 7 to be colored")
	}
	if !cm[19] || !cm[20] || !cm[21] {
		t.Errorf("expected 'kit' starting at pos 19 to be colored")
	}
	if cm[0] {
		t.Errorf("expected 'a' at pos 0 NOT to be colored")
	}
}

func TestBuildColorMap_Overlap(t *testing.T) {
	cm := buildColorMap("aaa", "aa")
	if !cm[0] {
		t.Errorf("expected position 0 to be colored")
	}
	if !cm[1] {
		t.Errorf("expected position 1 to be colored")
	}
	if !cm[2] {
		t.Errorf("expected position 2 to be colored")
	}
}

func TestRenderArtColor_EmptyInput(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArtColor("", banner, "\033[31m", "")
	if got != "" {
		t.Errorf("expected empty string, got %q", got)
	}
}

func TestRenderArtColor_WholeString(t *testing.T) {
	banner := loadStdTestBanner(t)
	result := renderArtColor("A", banner, "\033[31m", "")
	lines := strings.Split(strings.TrimSuffix(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(lines))
	}
	for _, line := range lines {
		if !strings.HasPrefix(line, "\033[31m") {
			t.Errorf("expected line to start with color code: %q", line)
		}
		if !strings.HasSuffix(line, "\033[0m") {
			t.Errorf("expected line to end with reset code: %q", line)
		}
	}
}

func TestRenderArtColor_Substring(t *testing.T) {
	banner := loadStdTestBanner(t)
	result := renderArtColor("hello", banner, "\033[31m", "ll")
	lines := strings.Split(strings.TrimSuffix(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(lines))
	}
	for _, line := range lines {
		if !strings.Contains(line, "\033[31m") {
			t.Errorf("expected line to contain color code: %q", line)
		}
		if !strings.Contains(line, "\033[0m") {
			t.Errorf("expected line to contain reset code: %q", line)
		}
	}
}

func TestRenderArtColor_NoColorWithoutSubstr(t *testing.T) {
	banner := loadStdTestBanner(t)
	plain := renderArt("hello", banner)
	colored := renderArtColor("hello", banner, "\033[31m", "zzz")
	if plain != colored {
		t.Errorf("expected no color when substr doesn't match")
	}
}

func TestRenderArtColor_SameAsPlainWhenSubstrEmptyAndNoColor(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArtColor("Hello", banner, "\033[0m", "")
	if !strings.HasPrefix(got, "\033[0m") || !strings.HasSuffix(strings.TrimRight(got, "\n"), "\033[0m") {
		t.Errorf("expected reset-colored output to have color/reset wrappers")
	}
}

func TestRenderArtColor_Multiline(t *testing.T) {
	banner := loadStdTestBanner(t)
	result := renderArtColor("a\\nb", banner, "\033[31m", "")
	n := strings.Count(result, "\n")
	if n != 16 {
		t.Errorf("expected 16 newlines for 2 lines, got %d", n)
	}
}

func TestParseColor_Black(t *testing.T) {
	got := parseColor("black")
	want := "\033[30m"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestParseColor_Green(t *testing.T) {
	got := parseColor("green")
	want := "\033[32m"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestParseColor_Yellow(t *testing.T) {
	got := parseColor("yellow")
	want := "\033[33m"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestParseColor_Magenta(t *testing.T) {
	got := parseColor("magenta")
	want := "\033[35m"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestParseColor_Cyan(t *testing.T) {
	got := parseColor("cyan")
	want := "\033[36m"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestParseColor_White(t *testing.T) {
	got := parseColor("white")
	want := "\033[37m"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestParseArgs_NoArgs(t *testing.T) {
	_, _, _, _, _, _, ok := parseArgs([]string{})
	if ok {
		t.Errorf("expected false for empty args")
	}
}

func TestParseArgs_JustString(t *testing.T) {
	color, substr, _, input, banner, _, ok := parseArgs([]string{"hello"})
	if !ok || color != "" || substr != "" || input != "hello" || banner != "standard" {
		t.Errorf("got color=%q substr=%q input=%q banner=%q ok=%v", color, substr, input, banner, ok)
	}
}

func TestParseArgs_StringAndBanner(t *testing.T) {
	color, substr, _, input, banner, _, ok := parseArgs([]string{"hello", "shadow"})
	if !ok || color != "" || substr != "" || input != "hello" || banner != "shadow" {
		t.Errorf("got color=%q substr=%q input=%q banner=%q ok=%v", color, substr, input, banner, ok)
	}
}

func TestParseArgs_ColorAndString(t *testing.T) {
	color, substr, _, input, banner, _, ok := parseArgs([]string{"--color=red", "hello"})
	if !ok || color != "red" || substr != "" || input != "hello" || banner != "standard" {
		t.Errorf("got color=%q substr=%q input=%q banner=%q ok=%v", color, substr, input, banner, ok)
	}
}

func TestParseArgs_ColorSubstrAndString(t *testing.T) {
	color, substr, _, input, banner, _, ok := parseArgs([]string{"--color=red", "ll", "hello"})
	if !ok || color != "red" || substr != "ll" || input != "hello" || banner != "standard" {
		t.Errorf("got color=%q substr=%q input=%q banner=%q ok=%v", color, substr, input, banner, ok)
	}
}

func TestParseArgs_ColorSubstrStringAndBanner(t *testing.T) {
	color, substr, _, input, banner, _, ok := parseArgs([]string{"--color=red", "ll", "hello", "thinkertoy"})
	if !ok || color != "red" || substr != "ll" || input != "hello" || banner != "thinkertoy" {
		t.Errorf("got color=%q substr=%q input=%q banner=%q ok=%v", color, substr, input, banner, ok)
	}
}

func TestParseArgs_ColorStringAndBanner(t *testing.T) {
	color, substr, _, input, banner, _, ok := parseArgs([]string{"--color=red", "hello", "shadow"})
	if !ok || color != "red" || substr != "" || input != "hello" || banner != "shadow" {
		t.Errorf("got color=%q substr=%q input=%q banner=%q ok=%v", color, substr, input, banner, ok)
	}
}

func TestParseArgs_InvalidColorFlag(t *testing.T) {
	_, _, _, _, _, _, ok := parseArgs([]string{"--color", "hello"})
	if ok {
		t.Errorf("expected false for '--color' without =")
	}
}

func TestParseArgs_EmptyColorValue(t *testing.T) {
	_, _, _, _, _, _, ok := parseArgs([]string{"--color=", "hello"})
	if ok {
		t.Errorf("expected false for empty color value")
	}
}

func TestParseArgs_TooManyArgs(t *testing.T) {
	_, _, _, _, _, _, ok := parseArgs([]string{"--color=red", "a", "b", "c", "d"})
	if ok {
		t.Errorf("expected false for too many args")
	}
}

func TestRenderArtColor_OffsetNewlines(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArtColor("\\n", banner, "\033[31m", "")
	want := "\n"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestRenderArtColor_KitExample(t *testing.T) {
	banner := loadStdTestBanner(t)
	result := renderArtColor("a king kitten have kit", banner, "\033[31m", "kit")
	lines := strings.Split(strings.TrimSuffix(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(lines))
	}
	for _, line := range lines {
		if !strings.Contains(line, "\033[31m") {
			t.Errorf("expected line to contain color code")
		}
		if !strings.Contains(line, "\033[0m") {
			t.Errorf("expected line to contain reset code")
		}
	}
}

func TestParseArgs_InvalidBannerNotPopped(t *testing.T) {
	_, _, _, _, _, _, ok := parseArgs([]string{"hello", "shadowy"})
	if ok {
		t.Errorf("expected false for unknown banner 'shadowy'")
	}
}

func TestParseArgs_StringWithValidBanner(t *testing.T) {
	color, substr, _, input, banner, _, ok := parseArgs([]string{"hello", "shadow"})
	if !ok || color != "" || substr != "" || input != "hello" || banner != "shadow" {
		t.Errorf("got color=%q substr=%q input=%q banner=%q ok=%v", color, substr, input, banner, ok)
	}
}

func TestParseArgs_AlignFlag(t *testing.T) {
	_, _, align, input, banner, _, ok := parseArgs([]string{"--align=center", "hello"})
	if !ok || align != "center" || input != "hello" || banner != "standard" {
		t.Errorf("got align=%q input=%q banner=%q ok=%v", align, input, banner, ok)
	}
}

func TestParseArgs_AlignFlagWithBanner(t *testing.T) {
	_, _, align, input, banner, _, ok := parseArgs([]string{"--align=right", "hello", "shadow"})
	if !ok || align != "right" || input != "hello" || banner != "shadow" {
		t.Errorf("got align=%q input=%q banner=%q ok=%v", align, input, banner, ok)
	}
}

func TestParseArgs_AlignAndColorFlags(t *testing.T) {
	_, _, align, input, banner, _, ok := parseArgs([]string{"--align=center", "--color=red", "hello"})
	if !ok || align != "center" || input != "hello" || banner != "standard" {
		t.Errorf("got align=%q input=%q banner=%q ok=%v", align, input, banner, ok)
	}
}

func TestParseArgs_ColorAndAlignFlags(t *testing.T) {
	_, _, align, input, banner, _, ok := parseArgs([]string{"--color=red", "--align=center", "hello"})
	if !ok || align != "center" || input != "hello" || banner != "standard" {
		t.Errorf("got align=%q input=%q banner=%q ok=%v", align, input, banner, ok)
	}
}

func TestParseArgs_InvalidAlignFlagFormat(t *testing.T) {
	_, _, _, _, _, _, ok := parseArgs([]string{"--align", "hello"})
	if ok {
		t.Errorf("expected false for '--align' without =")
	}
}

func TestParseArgs_OutputFlag(t *testing.T) {
	color, substr, _, input, banner, outputFile, ok := parseArgs([]string{"--output=out.txt", "hello", "standard"})
	if !ok || outputFile != "out.txt" || input != "hello" || banner != "standard" || color != "" || substr != "" {
		t.Errorf("got outputFile=%q input=%q banner=%q color=%q substr=%q ok=%v", outputFile, input, banner, color, substr, ok)
	}
}

func TestParseArgs_OutputFlagJustString(t *testing.T) {
	_, _, _, input, banner, outputFile, ok := parseArgs([]string{"--output=out.txt", "hello"})
	if !ok || outputFile != "out.txt" || input != "hello" || banner != "standard" {
		t.Errorf("got outputFile=%q input=%q banner=%q ok=%v", outputFile, input, banner, ok)
	}
}

func TestParseArgs_OutputFlagWithBanner(t *testing.T) {
	_, _, _, input, banner, outputFile, ok := parseArgs([]string{"--output=out.txt", "hello", "shadow"})
	if !ok || outputFile != "out.txt" || input != "hello" || banner != "shadow" {
		t.Errorf("got outputFile=%q input=%q banner=%q ok=%v", outputFile, input, banner, ok)
	}
}

func TestParseArgs_OutputFlagWithColor(t *testing.T) {
	color, _, _, input, banner, outputFile, ok := parseArgs([]string{"--output=out.txt", "--color=red", "hello"})
	if !ok || outputFile != "out.txt" || color != "red" || input != "hello" || banner != "standard" {
		t.Errorf("got outputFile=%q color=%q input=%q banner=%q ok=%v", outputFile, color, input, banner, ok)
	}
}

func TestParseArgs_OutputFlagWithAlign(t *testing.T) {
	_, _, align, input, banner, outputFile, ok := parseArgs([]string{"--output=out.txt", "--align=center", "hello"})
	if !ok || outputFile != "out.txt" || align != "center" || input != "hello" || banner != "standard" {
		t.Errorf("got outputFile=%q align=%q input=%q banner=%q ok=%v", outputFile, align, input, banner, ok)
	}
}

func TestParseArgs_OutputFlagInvalidNoEqual(t *testing.T) {
	_, _, _, _, _, _, ok := parseArgs([]string{"--output", "hello"})
	if ok {
		t.Errorf("expected false for '--output' without =")
	}
}

func TestParseArgs_OutputFlagEmptyValue(t *testing.T) {
	_, _, _, _, _, _, ok := parseArgs([]string{"--output=", "hello"})
	if ok {
		t.Errorf("expected false for empty output value")
	}
}

func TestParseArgs_OutputFlagDuplicate(t *testing.T) {
	_, _, _, _, _, _, ok := parseArgs([]string{"--output=a.txt", "--output=b.txt", "hello"})
	if ok {
		t.Errorf("expected false for duplicate --output flags")
	}
}

func TestCharWidth(t *testing.T) {
	banner := loadStdTestBanner(t)
	w := charWidth(0, banner)
	if w <= 0 {
		t.Errorf("expected positive width for space, got %d", w)
	}
}

func TestTextPixelWidth(t *testing.T) {
	banner := loadStdTestBanner(t)
	w := textPixelWidth("A", banner)
	if w <= 0 {
		t.Errorf("expected positive width for 'A', got %d", w)
	}
}

func TestRenderSegmentRight(t *testing.T) {
	banner := loadStdTestBanner(t)
	termWidth := 80
	result := renderSegmentRight("A", banner, termWidth)
	lines := strings.Split(strings.TrimSuffix(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(lines))
	}
	for _, line := range lines {
		if len(line) > termWidth {
			t.Errorf("line length %d exceeds terminal width %d", len(line), termWidth)
		}
	}
}

func TestRenderSegmentCenter(t *testing.T) {
	banner := loadStdTestBanner(t)
	termWidth := 80
	result := renderSegmentCenter("A", banner, termWidth)
	lines := strings.Split(strings.TrimSuffix(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(lines))
	}
	for _, line := range lines {
		if len(line) > termWidth {
			t.Errorf("line length %d exceeds terminal width %d", len(line), termWidth)
		}
	}
}

func TestRenderSegmentJustify(t *testing.T) {
	banner := loadStdTestBanner(t)
	termWidth := 80
	result := renderSegmentJustify("A B", banner, termWidth)
	lines := strings.Split(strings.TrimSuffix(result, "\n"), "\n")
	if len(lines) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(lines))
	}
	for _, line := range lines {
		if len(line) > termWidth {
			t.Errorf("line length %d exceeds terminal width %d", len(line), termWidth)
		}
	}
}

func TestRenderSegmentJustifySingleWord(t *testing.T) {
	banner := loadStdTestBanner(t)
	justify := renderSegmentJustify("Hello", banner, 80)
	left := renderSegmentLeft("Hello", banner)
	if justify != left {
		t.Errorf("justify with single word should match left alignment")
	}
}

func TestRenderArtAligned_DefaultIsLeft(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArtAligned("Hello", banner, "", 80)
	want := renderArt("Hello", banner)
	if got != want {
		t.Errorf("default align should produce left-aligned output")
	}
}

func TestRenderArtAligned_Left(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArtAligned("Hello", banner, "left", 80)
	want := renderArt("Hello", banner)
	if got != want {
		t.Errorf("left align should match renderArt")
	}
}

func TestRenderArtAligned_Empty(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArtAligned("", banner, "center", 80)
	if got != "" {
		t.Errorf("expected empty result for empty input, got %q", got)
	}
}

func TestRenderArtAligned_Newline(t *testing.T) {
	banner := loadStdTestBanner(t)
	got := renderArtAligned("\\n", banner, "right", 80)
	want := "\n"
	if got != want {
		t.Errorf("expected newline for newline input, got %q", got)
	}
}

func TestGetTerminalWidth(t *testing.T) {
	w := getTerminalWidth()
	if w <= 0 {
		t.Errorf("expected positive terminal width, got %d", w)
	}
}

func loadAllBannersForTest(t *testing.T) [][95][8]string {
	t.Helper()
	var banners [][95][8]string
	for _, name := range []string{"standard", "shadow", "thinkertoy"} {
		data, err := os.ReadFile(name + ".txt")
		if err != nil {
			t.Fatalf("failed to load %s: %v", name, err)
		}
		banners = append(banners, parseBanner(string(data)))
	}
	return banners
}

func TestReverseArt_RoundTripStandard(t *testing.T) {
	banners := loadAllBannersForTest(t)
	inputs := []string{"hello", "Hello World", "42", "{Hello!}", "abc123", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyz"}
	for _, input := range inputs {
		art := renderArt(input, banners[0])
		result := reverseArt(art, banners...)
		if result != input {
			t.Errorf("round-trip failed for %q: got %q", input, result)
		}
	}
}

func TestReverseArt_RoundTripShadow(t *testing.T) {
	banners := loadAllBannersForTest(t)
	inputs := []string{"hello", "Hello World", "42", "abc"}
	for _, input := range inputs {
		art := renderArt(input, banners[1])
		result := reverseArt(art, banners...)
		if result != input {
			t.Errorf("shadow round-trip failed for %q: got %q", input, result)
		}
	}
}

func TestReverseArt_RoundTripThinkertoy(t *testing.T) {
	banners := loadAllBannersForTest(t)
	inputs := []string{"hello", "Hello World", "42"}
	for _, input := range inputs {
		art := renderArt(input, banners[2])
		result := reverseArt(art, banners...)
		if result != input {
			t.Errorf("thinkertoy round-trip failed for %q: got %q", input, result)
		}
	}
}

func TestReverseArt_Multiline(t *testing.T) {
	banners := loadAllBannersForTest(t)
	inputs := []string{"hello\nthere", "a\nb\nc", "Hello\nWorld"}
	for _, input := range inputs {
		art := renderArt(input, banners[0])
		result := reverseArt(art, banners...)
		if result != input {
			t.Errorf("multiline round-trip failed for %q: got %q", input, result)
		}
	}
}

func TestReverseArt_EmptyData(t *testing.T) {
	banners := loadAllBannersForTest(t)
	result := reverseArt("", banners...)
	if result != "" {
		t.Errorf("expected empty for empty data, got %q", result)
	}

	art := renderArt("", banners[0])
	result = reverseArt(art, banners...)
	if result != "" {
		t.Errorf("expected empty for empty art, got %q", result)
	}
}

func TestReverseArt_WithNewline(t *testing.T) {
	banners := loadAllBannersForTest(t)
	input := "hello\nthere"
	art := renderArt(input, banners[0])
	result := reverseArt(art, banners...)
	if result != input {
		t.Errorf("round-trip failed for %q: got %q", input, result)
	}
}

func TestReverseArt_DetectsBanner(t *testing.T) {
	banners := loadAllBannersForTest(t)
	inputs := []struct {
		text   string
		shadow bool
	}{
		{"hello", false},
		{"Hello", false},
		{"Hello World", false},
	}
	for _, tc := range inputs {
		art := renderArt(tc.text, banners[0])
		result := reverseArt(art, banners...)
		if result != tc.text {
			t.Errorf("standard detection failed for %q: got %q", tc.text, result)
		}
	}
}

func TestReverseArt_FileNotFound(t *testing.T) {
	_, err := os.ReadFile("nonexistent.txt")
	if err == nil {
		t.Error("expected error for nonexistent file")
	}
}

func TestReverseArt_SpecialChars(t *testing.T) {
	banners := loadAllBannersForTest(t)
	input := "{Hello!} @#$%^&*()"
	art := renderArt(input, banners[0])
	result := reverseArt(art, banners...)
	if result != input {
		t.Errorf("special chars round-trip failed: got %q", result)
	}
}
