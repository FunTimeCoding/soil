package unit

import (
	"net"
	"os"
	"syscall"
)

func refusedConnection() error {
	return &net.OpError{
		Op:  "dial",
		Net: "tcp",
		Err: os.NewSyscallError("connect", syscall.ECONNREFUSED),
	}
}
