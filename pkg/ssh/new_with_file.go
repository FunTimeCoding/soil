package ssh

import (
	"golang.org/x/crypto/ssh"
	"io"
)

func NewWithFile(
	user string,
	host string,
	keyPath string,
	keyName string,
	secure bool,
) *Client {
	return &Client{
		user:   user,
		host:   host,
		secure: secure,
		authenticate: func() (ssh.AuthMethod, io.Closer, error) {
			m, e := fileAuthentication(keyPath, keyName)

			return m, nil, e
		},
		Panic: true,
	}
}
