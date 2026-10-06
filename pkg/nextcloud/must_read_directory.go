package nextcloud

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func (c *Client) MustReadDirectory(path string) []os.FileInfo {
	result, e := c.ReadDirectory(path)
	errors.PanicOnError(e)

	return result
}
