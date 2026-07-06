package author

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func (c *Client) ReadDirectory(path string) []os.FileInfo {
	result, e := c.client.ReadDir(path)
	errors.PanicOnError(e)

	return result
}
