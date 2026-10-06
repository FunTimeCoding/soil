package nextcloud

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func (c *Client) UploadFile(
	path string,
	sourcePath string,
	sourceName string,
) error {
	r, e := os.OpenRoot(sourcePath)

	if e != nil {
		return e
	}

	defer errors.LogClose(r)
	b, f := r.ReadFile(sourceName)

	if f != nil {
		return f
	}

	return c.author.WriteFile(path, b)
}
