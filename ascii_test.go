package main

import (
	"reflect"
	"testing"
)

// func Test_getLetterByCodeAscci(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			getLetterByCodeAscci()
// 		})
// 	}
// }

func Test_getLines(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []string
	}{
		{"test standard", "test.txt", []string{"yolo les kikis!", "yala les kakas?", "yili les kokos..."}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLines(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPosition(t *testing.T) {

	tests := []struct {
		name string
		in   byte
		want int
	}{
		{"test a", 'a', 587},
		{"test b", 'b', 596},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getPosition(tt.in)

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func Test_getCharLines(t *testing.T) {
	type args struct {
		char  byte
		lines []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test a", args{'a', getLines("standard.txt")}, []string{"        ", "        ", "  __ _  ", " / _` | ", "| (_| | ", " \\__,_| ", "        ", "        "}},
		{"test b", args{'b', getLines("standard.txt")}, []string{" _      ", "| |     ", "| |__   ", "| '_ \\  ", "| |_) | ", "|_.__/  ", "        ", "        "}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCharLines(tt.args.char, tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCharLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getStringLines(t *testing.T) {
	type args struct {
		s     string
		lines []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"test yolo", args{"y", getLines("standard.txt")}, [][]string{{"        ", "        ", " _   _  ", "| | | | ", "| |_| | ", " \\__, | ", " __/ /  ", "|___/   "}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStringLines(tt.args.s, tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStringLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
