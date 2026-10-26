package main

import "testing"

func Test_getColorCode(t *testing.T) {
	type args struct {
		color string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"black", args{"black"}, "\033[30m"},
		{"red", args{"red"}, "\033[31m"},
		{"green", args{"green"}, "\033[32m"},
		{"yellow", args{"yellow"}, "\033[33m"},
		{"blue", args{"blue"}, "\033[34m"},
		{"magenta", args{"magenta"}, "\033[35m"},
		{"cyan", args{"cyan"}, "\033[36m"},
		{"white", args{"white"}, "\033[37m"},
		{"invalid color", args{"invalid"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getColorCode(tt.args.color); got != tt.want {
				t.Errorf("getColorCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
