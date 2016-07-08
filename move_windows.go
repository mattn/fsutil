// +build windows

package fsutil

import (
	"syscall"
	"unsafe"
)

var (
	procMoveFile = kernel32.MustFindProc("MoveFileW")
)

func MoveFile(dst, src string) error {
	psrc, err := syscall.UTF16PtrFromString(src)
	if err != nil {
		return err
	}
	pdst, err := syscall.UTF16PtrFromString(dst)
	if err != nil {
		return err
	}
	r, _, err := procMoveFile.Call(uintptr(unsafe.Pointer(psrc)), uintptr(unsafe.Pointer(pdst)), uintptr(0))
	if r == 0 {
		return err
	}
	return nil
}
