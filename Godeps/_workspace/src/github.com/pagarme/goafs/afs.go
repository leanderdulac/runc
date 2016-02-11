package goafs

import (
	"unsafe"
	"github.com/paypal/gatt/linux/gioctl"
)

type ViceIoctl struct {
	In      uintptr
	Out     uintptr
	InSize  int16
	OutSize int16
}

// Creates a new PAG and attach it to the current process
func Setpag() error {
	return afs_syscall(21, 0, 0, 0, 0)
}

// Destroy all AFS tokens
func Unlog() error {
	iob := ViceIoctl{0, 0, 0, 0}

	return afs_pioctl(nil, gioctl.IoW(86, 9, unsafe.Sizeof(ViceIoctl{})), &iob, false)
}

func afs_pioctl(path *string, cmd uintptr, cmarg *ViceIoctl, follow bool) error {
	var followi uintptr

	if follow {
		followi = 1
	} else {
		followi = 0
	}

	return afs_syscall(20, uintptr(unsafe.Pointer(&path)), cmd, uintptr(unsafe.Pointer(cmarg)), followi)
}
