package server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"net"
	"path/filepath"
)

func (s *Server) Listen() error {
	system.MakeDirectory(filepath.Dir(s.socketPath))
	removeSocket(s.socketPath)
	listener, e := net.Listen("unix", s.socketPath)

	if e != nil {
		return e
	}

	defer errors.PanicClose(listener)
	defer removeSocket(s.socketPath)

	for {
		connection, e := listener.Accept()

		if e != nil {
			continue
		}

		go s.handleConnection(connection)
	}
}
