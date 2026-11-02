package ssh

import "github.com/pkg/sftp"

func (c *Client) sftpClient() (*sftp.Client, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.sftp != nil {
		return c.sftp, nil
	}

	d, e := c.dialed()

	if e != nil {
		return nil, e
	}

	result, f := sftp.NewClient(d)

	if f != nil {
		return nil, c.refused("open sftp on", c.host, f)
	}

	c.sftp = result

	return c.sftp, nil
}
