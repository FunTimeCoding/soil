package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/client"
	"net"
	"net/url"
	"os"
	"syscall"
	"testing"
)

func TestUnreachable(t *testing.T) {
	refused := &url.Error{
		Op:  "Get",
		URL: "http://host.example:8080/status",
		Err: &net.OpError{
			Op:  "dial",
			Net: "tcp",
			Err: os.NewSyscallError("connect", syscall.ECONNREFUSED),
		},
	}
	assert.True(t, client.Unreachable(refused))
	assert.True(
		t,
		client.Unreachable(
			&net.DNSError{
				Err:        "no such host",
				Name:       "host.example",
				IsNotFound: true,
			},
		),
	)
	assert.True(
		t,
		client.Unreachable(
			&url.Error{
				Op:  "Get",
				URL: "http://host.example:8080/status",
				Err: &net.OpError{
					Op:  "dial",
					Net: "tcp",
					Err: os.NewSyscallError("connect", syscall.EHOSTUNREACH),
				},
			},
		),
	)
	assert.False(t, client.Unreachable(errors.New("unexpected")))
}
