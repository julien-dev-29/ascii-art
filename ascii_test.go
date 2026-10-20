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

func Test_getALines(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getALines()
		})
	}
}
