package main

import "strings"

func containAntiSlashN(s string) bool {
	return strings.ContainsRune(s, '\n')
}
