package main

import (
	"reflect"
	"testing"
)

func Test_parseFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want map[rune][]string
	}{
		{"Test with a valid file", args{"test_banner.txt"}, map[rune][]string{
			' ': {
				"        ",
				"        ",
				"        ",
				"        ",
				"        ",
				"        ",
				"        ",
				"        ",
			},
			'!': {
				"   ##   ",
				"   ##   ",
				"   ##   ",
				"   ##   ",
				"   ##   ",
				"        ",
				"   ##   ",
				"        ",
			},
		},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseFile(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
