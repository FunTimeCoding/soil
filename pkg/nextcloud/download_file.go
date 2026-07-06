package nextcloud

import "github.com/funtimecoding/soil/pkg/system"

func (c *Client) DownloadFile(
	path string,
	destination string,
) {
	system.SaveBytes(destination, c.author.ReadFile(path))
}
