package gout

import (
	"syscall"
	"unsafe"
)

// console info and utilities

func ConsInfo() winsize {
	ws := winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&ws)))
	if int(retCode) == -1 {
		panic(errno)
	}
	return ws
}
