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
	result := New(user, host, secure)
	result.authenticate = func() (ssh.AuthMethod, io.Closer, error) {
		m, e := fileAuthentication(keyPath, keyName)

		return m, nil, e
	}

	return result
}
