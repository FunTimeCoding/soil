package nextcloud

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDownloadFile(
	path string,
	destination string,
) {
	errors.PanicOnError(c.DownloadFile(path, destination))
}
