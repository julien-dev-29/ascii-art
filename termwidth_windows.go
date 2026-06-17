//go:build windows

package main

import (
	"syscall"
	"unsafe"
)

func getTerminalWidth() int {
	mod := syscall.NewLazyDLL("kernel32.dll")
	proc := mod.NewProc("GetConsoleScreenBufferInfo")
	h := mod.NewProc("GetStdHandle")

	out, _, _ := h.Call(uintptr(^uint32(10) + 1)) // STD_OUTPUT_HANDLE = -11
	if out == 0 {
		return 80
	}

	var info struct {
		SizeX, SizeY                                     int16
		CursorX, CursorY                                 int16
		Attrs                                            uint16
		WindowLeft, WindowTop, WindowRight, WindowBottom int16
		MaxSizeX, MaxSizeY                               int16
	}
	ret, _, _ := proc.Call(out, uintptr(unsafe.Pointer(&info)))
	if ret == 0 {
		return 80
	}
	if w := int(info.SizeX); w > 0 {
		return w
	}
	return 80
}
