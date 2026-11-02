package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) ReadFile(path string) ([]byte, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	b, okay := c.files[path]

	if !okay {
		return nil, not_found.Format("read %s: file does not exist", path)
	}

	return b, nil
}
