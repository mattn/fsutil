// +build windows

package fsutil

import "syscall"

var (
	kernel32 = syscall.MustLoadDLL("kernel32.dll")
)
