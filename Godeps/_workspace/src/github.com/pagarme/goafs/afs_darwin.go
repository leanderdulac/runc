// +build darwin

package goafs

import (
	"github.com/paypal/gatt/linux/gioctl"
	"syscall"
	"unsafe"
)

const intSize = 32 + int(^uintptr(0)>>63<<5)

type afsprocdata32 struct {
	syscall uintptr
	param1 uintptr
	param2 uintptr
	param3 uintptr
	param4 uintptr
	param5 uintptr
	param6 uintptr
	retval uintptr
}

type afsprocdata64 struct {
	param1 uintptr
	param2 uintptr
	param3 uintptr
	param4 uintptr
	param5 uintptr
	param6 uintptr
	syscall int32
	retval int32
}

func afs_syscall(call uintptr, param1 uintptr, param2 uintptr, param3 uintptr, param4 uintptr) error {
	fd, err := syscall.Open("/dev/openafs_ioctl", syscall.O_RDWR, 0)

	if err != nil {
		return err
	}

	if intSize == 32 {
		data := afsprocdata32{
			call,
			param1,
			param2,
			param3,
			param4,
			0,
			0,
			0,
		}

		err = gioctl.Ioctl(uintptr(fd), gioctl.IoW(67, 1, unsafe.Sizeof(uintptr(0))), uintptr(unsafe.Pointer(&data)))
	} else {
		data := afsprocdata64{
			param1,
			param2,
			param3,
			param4,
			0,
			0,
			int32(call),
			0,
		}

		err = gioctl.Ioctl(uintptr(fd), gioctl.IoW(67, 2, unsafe.Sizeof(uintptr(0))), uintptr(unsafe.Pointer(&data)))
	}

	syscall.Close(fd)

	return err
}

