package nextcloud

import "github.com/funtimecoding/soil/pkg/system"

func (c *Client) UploadFile(
	path string,
	sourcePath string,
	sourceName string,
) {
	c.author.WriteFile(path, system.ReadBytes(sourcePath, sourceName))
}
