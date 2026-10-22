package main

import "testing"

func Test_containAntiSlashN(t *testing.T) {

	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"with \n", "yolo les \n kikis", true},
		{"without \n", "yolo les kikis", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containAntiSlashN(tt.in); got != tt.want {
				t.Errorf("containAntiSlashN() = %v, want %v", got, tt.want)
			}
		})
	}
}
