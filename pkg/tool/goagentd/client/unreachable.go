package client

import (
	"errors"
	"net"
	"syscall"
)

func Unreachable(e error) bool {
	var d *net.DNSError

	if errors.As(e, &d) {
		return true
	}

	if errors.Is(e, syscall.ECONNREFUSED) ||
		errors.Is(e, syscall.EHOSTUNREACH) ||
		errors.Is(e, syscall.ENETUNREACH) ||
		errors.Is(e, syscall.ETIMEDOUT) {
		return true
	}

	var n net.Error

	return errors.As(e, &n) && n.Timeout()
}
