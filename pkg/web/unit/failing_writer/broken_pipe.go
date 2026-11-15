package failing_writer

import (
	"net"
	"os"
	"syscall"
)

func BrokenPipe() error {
	return &net.OpError{
		Op:  "write",
		Net: "tcp",
		Err: os.NewSyscallError("write", syscall.EPIPE),
	}
}
