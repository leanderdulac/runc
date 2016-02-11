// +build linux

package goafs

import (
	"unsafe"
	"syscall"

	"github.com/paypal/gatt/linux/gioctl"
)

type afsprocdata struct {
	param4  uintptr
	param3  uintptr
	param2  uintptr
	param1  uintptr
	syscall uintptr
}

func afs_syscall(call uintptr, param1 uintptr, param2 uintptr, param3 uintptr, param4 uintptr) error {
	fd, err := syscall.Open("/proc/fs/openafs/afs_ioctl", syscall.O_RDWR, 0)

	if err != nil {
		fd, err = syscall.Open("/proc/fs/nnpfs/afs_ioctl", syscall.O_RDWR, 0)
	}

	if err != nil {
		return err
	}

	data := afsprocdata{
		param4,
		param3,
		param2,
		param1,
		call,
	}

	err = gioctl.Ioctl(uintptr(fd), gioctl.IoW(67, 1, unsafe.Sizeof(uintptr(0))), uintptr(unsafe.Pointer(&data)))

	syscall.Close(fd)

	return err
}
