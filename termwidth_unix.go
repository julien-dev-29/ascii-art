//go:build !windows

package main

import (
	"os"
	"syscall"
	"unsafe"
)

func getTerminalWidth() int {
	fd := os.Stdout.Fd()
	ws := struct {
		Row    uint16
		Col    uint16
		Xpixel uint16
		Ypixel uint16
	}{}
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&ws)))
	if err != 0 {
		return 80
	}
	if ws.Col == 0 {
		return 80
	}
	return int(ws.Col)
}
