package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) RemoveFile(path string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e := c.failures.Caller(); e != nil {
		return e
	}

	if _, okay := c.files[path]; !okay {
		return not_found.Format("remove %s: file does not exist", path)
	}

	delete(c.files, path)

	return nil
}
