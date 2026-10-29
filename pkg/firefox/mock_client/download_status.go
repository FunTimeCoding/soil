package mock_client

import "github.com/funtimecoding/soil/pkg/firefox/download"

func (c *Client) DownloadStatus(identifier int) ([]*download.Download, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	if identifier <= 0 {
		return c.downloads, nil
	}

	var result []*download.Download

	for _, d := range c.downloads {
		if d.Identifier == identifier {
			result = append(result, d)
		}
	}

	return result, nil
}
