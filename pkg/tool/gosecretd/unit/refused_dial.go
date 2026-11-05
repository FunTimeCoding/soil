package unit

import (
	"net"
	"os"
	"syscall"
)

func refusedDial() error {
	return &net.OpError{
		Op:  "dial",
		Net: "tcp",
		Err: os.NewSyscallError("connect", syscall.ECONNREFUSED),
	}
}
