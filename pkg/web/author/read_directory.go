package author

import "os"

func (c *Client) ReadDirectory(path string) ([]os.FileInfo, error) {
	return c.client.ReadDir(path)
}
