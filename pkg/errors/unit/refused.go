package unit

import (
	"net"
	"net/url"
	"os"
	"syscall"
)

func refused() error {
	return &url.Error{
		Op:  "Get",
		URL: "http://alfa.example:8080/api/items?api_key=secret",
		Err: &net.OpError{
			Op:  "dial",
			Net: "tcp",
			Err: os.NewSyscallError("connect", syscall.ECONNREFUSED),
		},
	}
}
