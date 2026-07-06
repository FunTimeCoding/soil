package systemd

import "github.com/funtimecoding/soil/pkg/ssh"

func New(c *ssh.Client) *Client {
	return &Client{ssh: c}
}
