package server_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"golang.org/x/crypto/ssh"
)

func (s *Server) serveSession(
	c ssh.Channel,
	requests <-chan *ssh.Request,
) {
	defer s.group.Done()

	for r := range requests {
		if r.Type != "subsystem" || string(r.Payload[4:]) != "sftp" {
			errors.LogOnError(r.Reply(false, nil))

			continue
		}

		errors.LogOnError(r.Reply(true, nil))
		serveSftp(c)

		return
	}

	release(c)
}
