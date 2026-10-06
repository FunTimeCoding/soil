package connection

import (
	"context"
	"crypto/x509"
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"io"
	"net"
	"strings"
	"syscall"
)

func kind(e error) (string, string) {
	var network net.Error
	var name *net.DNSError
	var authority x509.UnknownAuthorityError
	var invalid x509.CertificateInvalidError
	var hostname x509.HostnameError

	switch {
	case errors.Is(e, context.DeadlineExceeded),
		errors.As(e, &network) && network.Timeout():

		return constant.Timeout, constant.TimedOut
	case errors.As(e, &name):
		return constant.Unreachable, constant.UnknownHost
	case errors.As(e, &authority),
		errors.As(e, &invalid),
		errors.As(e, &hostname):

		return constant.Unreachable, constant.Untrusted
	case errors.Is(e, syscall.ECONNREFUSED):
		return constant.Unreachable, constant.Refused
	case errors.Is(e, syscall.EHOSTUNREACH),
		errors.Is(e, syscall.ENETUNREACH):

		return constant.Unreachable, constant.NoRoute
	case errors.Is(e, io.ErrUnexpectedEOF),
		errors.Is(e, io.EOF),
		errors.Is(e, syscall.ECONNRESET),
		errors.Is(e, syscall.EPIPE),
		strings.Contains(e.Error(), constant.StreamClosed):

		return constant.Dropped, constant.Broken
	default:
		return "", ""
	}
}
