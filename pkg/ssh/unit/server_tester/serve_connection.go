package server_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"golang.org/x/crypto/ssh"
	"net"
)

func (s *Server) serveConnection(c net.Conn) {
	defer s.group.Done()
	_, channels, requests, e := ssh.NewServerConn(c, s.configuration)

	if e != nil {
		errors.LogClose(c)

		return
	}

	go ssh.DiscardRequests(requests)

	for n := range channels {
		if n.ChannelType() != "session" {
			errors.LogOnError(n.Reject(ssh.UnknownChannelType, "session only"))

			continue
		}

		channel, channelRequests, f := n.Accept()

		if f != nil {
			continue
		}

		s.group.Add(1)
		go s.serveSession(channel, channelRequests)
	}
}
