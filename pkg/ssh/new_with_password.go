package ssh

import "golang.org/x/crypto/ssh"

func NewWithPassword(
	user string,
	host string,
	password string,
	secure bool,
) *Client {
	return &Client{
		user:   user,
		host:   host,
		secure: secure,
		authenticate: func() (ssh.AuthMethod, error) {
			return ssh.Password(password), nil
		},
		Panic: true,
	}
}
