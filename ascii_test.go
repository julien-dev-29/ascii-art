package main

import (
	"os"
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

func Test_splitByLine(t *testing.T) {
	testData, _ := os.ReadFile("test.txt")
	tests := []struct {
		name string
		in   []byte
		want []string
	}{
		{"test standard", testData, []string{"yolo\r\n", "les\r\n", "kikis!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByLine(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
