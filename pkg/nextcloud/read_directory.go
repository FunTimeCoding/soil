package nextcloud

import "os"

func (c *Client) ReadDirectory(path string) ([]os.FileInfo, error) {
	return c.author.ReadDirectory(path)
}
