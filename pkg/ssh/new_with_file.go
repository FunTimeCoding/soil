package ssh

import "golang.org/x/crypto/ssh"

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
		authenticate: func() (ssh.AuthMethod, error) {
			return fileAuthentication(keyPath, keyName)
		},
		Panic: true,
	}
}
