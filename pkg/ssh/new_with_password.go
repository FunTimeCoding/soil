package ssh

import (
	"golang.org/x/crypto/ssh"
	"io"
)

func NewWithPassword(
	user string,
	host string,
	password string,
	secure bool,
) *Client {
	result := New(user, host, secure)
	result.authenticate = func() (ssh.AuthMethod, io.Closer, error) {
		return ssh.Password(password), nil, nil
	}

	return result
}
