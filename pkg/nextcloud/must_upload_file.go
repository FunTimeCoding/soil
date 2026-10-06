package nextcloud

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustUploadFile(
	path string,
	sourcePath string,
	sourceName string,
) {
	errors.PanicOnError(c.UploadFile(path, sourcePath, sourceName))
}
