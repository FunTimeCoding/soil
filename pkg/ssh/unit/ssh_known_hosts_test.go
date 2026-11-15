package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/ssh"
	"github.com/funtimecoding/soil/pkg/ssh/unit/server_tester"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"golang.org/x/crypto/ssh/knownhosts"
	"net"
	"strconv"
	"testing"
)

func TestSecureDialNegotiatesKeyTypeKnownHostsHolds(t *testing.T) {
	s := server_tester.New(t)
	system_tester.Write(
		t,
		system_tester.Home(t),
		".ssh/known_hosts",
		knownhosts.Line(
			[]string{
				knownhosts.Normalize(
					net.JoinHostPort("127.0.0.1", strconv.Itoa(s.Port)),
				),
			},
			s.HostKey,
		),
	)
	c := ssh.NewWithPassword("alfa", "127.0.0.1", s.Password, true).
		WithPort(s.Port)
	defer c.Close()
	assert.FatalOnError(t, c.Dial())
}

func TestSecureDialRefusesUnknownHost(t *testing.T) {
	s := server_tester.New(t)
	system_tester.Write(t, system_tester.Home(t), ".ssh/known_hosts", "")
	c := ssh.NewWithPassword("alfa", "127.0.0.1", s.Password, true).
		WithPort(s.Port)
	defer c.Close()
	assert.NotNil(t, c.Dial())
}
