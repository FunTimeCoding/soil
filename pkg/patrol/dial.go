package patrol

import "github.com/funtimecoding/soil/pkg/ssh"

func (c *Client) dial() (*ssh.Client, error) {
	result := ssh.NewWithFile(
		c.user,
		c.host,
		c.keyDirectory,
		c.keyName,
		true,
	).NoPanic()

	if e := result.Dial(); e != nil {
		return nil, e
	}

	return result, nil
}
