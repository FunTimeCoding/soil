package unit

import (
	"net"
	"os"
	"syscall"
)

func brokenPipe() error {
	return &net.OpError{
		Op:  "write",
		Net: "tcp",
		Err: os.NewSyscallError("write", syscall.EPIPE),
	}
}
