package proxmox

func WithUntrusted() Option {
	return func(c *Client) {
		c.untrusted = true
	}
}
