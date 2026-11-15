package outage_tester

import (
	"net"
	"os"
	"syscall"
)

func RefusedDial() error {
	return &net.OpError{
		Op:  "dial",
		Net: "tcp",
		Err: os.NewSyscallError("connect", syscall.ECONNREFUSED),
	}
}
