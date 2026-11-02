package server_tester

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

func (s *Server) checkPassword(
	m ssh.ConnMetadata,
	password []byte,
) (*ssh.Permissions, error) {
	if string(password) != s.Password {
		return nil, fmt.Errorf("password refused for %s", m.User())
	}

	return nil, nil
}
